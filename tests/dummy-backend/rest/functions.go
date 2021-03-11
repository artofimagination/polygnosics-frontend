package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
		return
	}
	fmt.Println("Failed to pretty print data")
}

func writeResponse(data string, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, data)
}

func decodeRequest(r *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func encodeResponse(data map[string]interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(data)
	if err != nil {
		err = errors.Wrap(errors.WithStack(err), "Failed to encode response")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}
	writeResponse(string(b), w, http.StatusOK)
}

func (c *Controller) updateTestData(w http.ResponseWriter, r *http.Request) {
	requestData, err := decodeRequest(r)
	if err != nil {
		err = errors.Wrap(errors.WithStack(err), "Failed to decode request")
		writeResponse(fmt.Sprintf("{'error':'%s'}", err.Error()), w, http.StatusInternalServerError)
		return
	}

	c.TestData = requestData
}

func uploadFile(key string, fileName string, r *http.Request) error {
	file, handler, err := r.FormFile(key)
	if err == http.ErrMissingFile {
		return nil
	}

	if err != nil {
		return err
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	defer file.Close()

	// Create file
	dst, err := os.Create(fmt.Sprintf("/user-assets/uploads/%s", fileName))
	if err != nil {
		return err
	}

	// Copy the uploaded file to the created file on the file system.
	if _, err := io.Copy(dst, file); err != nil {
		if err2 := dst.Close(); err2 != nil {
			err = errors.Wrap(errors.WithStack(err), err2.Error())
		}
		return err
	}
	dst.Close()

	return nil
}
