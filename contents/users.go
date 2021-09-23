package contents

import (
	"fmt"
	"net/http"

	"github.com/artofimagination/polygnosics-frontend/rest/backend"
)

// Details and assets field keys
const (
	UserMapKey                 = "user"
	UserAvatarKey              = "avatar"
	UserProfilePathKey         = "profile"
	UserProfileEditPathKey     = "profile_edit"
	UserProfileAvatarUploadKey = "avatar_upload"
	UserNameKey                = "username"
	UserEmailKey               = "email"
	UserFullNameKey            = "full_name"
	UserDataMapKey             = "datamap"
	UserCountryKey             = "country"
	UserCityKey                = "city"
	UserWebsiteKey             = "website"
	UserPhoneKey               = "phone"
	UserAboutKey               = "about"
	UserFacebookKey            = "facebook_link"
	UserTwitterKey             = "twitter_link"
	UserGithubKey              = "github_link"
	UserProductWizardPathKey   = "wizard"

	UserLocationKey = "location"
)

// GetUserContent fills a string nested map with all user details and assets info
func (c *ContentController) GetUserContent(user *backend.User) map[string]interface{} {
	content := make(map[string]interface{})
	content[FutureFeature] = 1
	content[CreateItemKey] = ""
	content[UserMapKey] = make(map[string]interface{})
	userData := content[UserMapKey].(map[string]interface{})
	userData[UserNameKey] = user.UserName
	userData[UserEmailKey] = user.Email
	for k, v := range user.Settings[UserDataMapKey].(map[string]interface{}) {
		userData[k] = v
	}
	for k, v := range user.Assets[UserDataMapKey].(map[string]interface{}) {
		userData[k] = v
	}

	userData[UserLocationKey] =
		setLocationString(
			userData[UserCountryKey].(string),
			userData[UserCityKey].(string))

	userData[UserProfileAvatarUploadKey] = "Upload your avatar"
	userData[UserProfilePathKey] = fmt.Sprintf("/user-main/profile?user=%s", user.ID)
	userData[UserProfileEditPathKey] = "/user-main/profile-edit"
	userData[UserProductWizardPathKey] = fmt.Sprintf("/user-main/product-wizard?user=%s", user.ID)
	return content
}

func (c *ContentController) StoreUserInfo(r *http.Request) {
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserNameKey] = r.FormValue(UserNameKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserFullNameKey] = r.FormValue(UserFullNameKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserCountryKey] = r.FormValue(UserCountryKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserCityKey] = r.FormValue(UserCityKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserPhoneKey] = r.FormValue(UserPhoneKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserWebsiteKey] = r.FormValue(UserWebsiteKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserAboutKey] = r.FormValue(UserAboutKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserFacebookKey] = r.FormValue(UserFacebookKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserTwitterKey] = r.FormValue(UserTwitterKey)
	c.User.Settings[UserDataMapKey].(map[string]interface{})[UserGithubKey] = r.FormValue(UserGithubKey)
}
