package contents

import (
	"fmt"
	"net/http"
	"polygnosics-frontend/restbackend"
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
	UserCountryKey             = "country"
	UserCityKey                = "city"
	UserWebsiteKey             = "website"
	UserPhoneKey               = "phone"
	UserAboutKey               = "about"
	UserFacebookKey            = "facebook_link"
	UserTwitterKey             = "twitter_link"
	UserGithubKey              = "github_link"
	UserProductWizardPathKey   = "wizard"
)

// GetUserContent fills a string nested map with all user details and assets info
func (c *ContentController) GetUserContent(user *restbackend.User) map[string]interface{} {
	content := make(map[string]interface{})
	content[FutureFeature] = 1
	content[UserMapKey] = make(map[string]interface{})
	userData := content[UserMapKey].(map[string]interface{})
	userData[UserNameKey] = user.UserName
	userData[UserEmailKey] = user.Email
	for k, v := range user.Settings {
		userData[k] = v
	}
	for k, v := range user.Assets {
		userData[k] = v
	}
	userData[UserProfileAvatarUploadKey] = "Upload your avatar"
	userData[UserProfilePathKey] = fmt.Sprintf("/user-main/profile?user=%s", user.ID)
	userData[UserProfileEditPathKey] = "/user-main/profile-edit"
	userData[UserProductWizardPathKey] = fmt.Sprintf("/user-main/product-wizard?user=%s", user.ID)
	return content
}

func (c *ContentController) StoreUserInfo(r *http.Request) {
	c.User.Settings[UserNameKey] = r.FormValue(UserNameKey)
	c.User.Settings[UserFullNameKey] = r.FormValue(UserFullNameKey)
	c.User.Settings[UserCountryKey] = r.FormValue(UserCountryKey)
	c.User.Settings[UserCityKey] = r.FormValue(UserCityKey)
	c.User.Settings[UserPhoneKey] = r.FormValue(UserPhoneKey)
	c.User.Settings[UserWebsiteKey] = r.FormValue(UserWebsiteKey)
	c.User.Settings[UserAboutKey] = r.FormValue(UserAboutKey)
	c.User.Settings[UserFacebookKey] = r.FormValue(UserFacebookKey)
	c.User.Settings[UserTwitterKey] = r.FormValue(UserTwitterKey)
	c.User.Settings[UserGithubKey] = r.FormValue(UserGithubKey)
}
