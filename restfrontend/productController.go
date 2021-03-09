package restfrontend

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *RESTFrontend) MyProducts(w http.ResponseWriter, r *http.Request) {
	content, err := c.ContentController.BuildMyProductsContent()
	if err != nil {
		errString := fmt.Sprintf("Failed to get product content. %s", errors.WithStack(err))
		c.RenderTemplate(w, UserMainMyProducts, c.ContentController.BuildErrorContent(errString))
		return
	}
	c.RenderTemplate(w, UserMainMyProducts, content)
}

func (c *RESTFrontend) MyProductDetails(w http.ResponseWriter, r *http.Request) {
	productID, err := parseItemID(r)
	if err != nil {
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent("Failed to parse product id"))
		return
	}

	content, err := c.ContentController.BuildProductDetailsContent(productID)
	if err != nil {
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent("Failed to get product content"))
		return
	}

	c.RenderTemplate(w, "details", content)
}

func (c *RESTFrontend) ProductDetails(w http.ResponseWriter, r *http.Request) {
	productID, err := parseItemID(r)
	if err != nil {
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent("Failed to parse product id"))
		return
	}

	content, err := c.ContentController.BuildProductDetailsContent(productID)
	if err != nil {
		c.RenderTemplate(w, UserMain, c.ContentController.BuildErrorContent("Failed to get product content"))
		return
	}

	c.RenderTemplate(w, "details", content)
}

func (c *RESTFrontend) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		content, err := c.ContentController.BuildProductWizardContent()
		if err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.RenderTemplate(w, "product-wizard", content)
	} else {
		content, err := c.ContentController.BuildUserMainContent()
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to load user main content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		if err := r.ParseMultipartForm(10 << 20); err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
			return
		}

		if err := c.RESTBackend.AddProduct(c.ContentController.User.ID, r); err != nil {
			c.HandleError(w, err.Error(), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.RenderTemplate(w, UserMain, content)
	}
}

func (c *RESTFrontend) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		productID, err := parseItemID(r)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to parse product id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		if err := c.RESTBackend.DeleteProduct(productID); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to delete product. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.MyProducts(w, r)
	}
}

func (c *RESTFrontend) EditProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := parseItemID(r)
	if err != nil {
		c.HandleError(w, fmt.Sprintf("Failed to parse product id. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
		return
	}

	if r.Method == GET {
		content, err := c.ContentController.BuildProductEditContent(productID)
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to get build product edit content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		c.RenderTemplate(w, "product-edit", content)
	} else {
		if err := c.RESTBackend.UpdateProduct(r); err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to update product. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}

		content, err := c.ContentController.BuildMyProductsContent()
		if err != nil {
			c.HandleError(w, fmt.Sprintf("Failed to get my products content. %s", errors.WithStack(err)), http.StatusInternalServerError, UserMainPath)
			return
		}
		c.RenderTemplate(w, UserMainMyProducts, content)
	}
}
