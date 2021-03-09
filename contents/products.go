package contents

import (
	"fmt"
	"polygnosics-frontend/restbackend"
)

// Details and assets field keys
const (
	ProductMapKey                = "product"
	ProductOwnerNameKey          = "owner_name"
	ProductOwnerPageNameKey      = "owner_page"
	ProductDetailPageKey         = "detail_path"
	ProductEditPageKey           = "edit_path"
	ProductPriceStringKey        = "price"
	ProductPublicTextKey         = "is_public_text"
	ProductRequires3DTextKey     = "requires_3d_text"
	ProductSupportsClientTextKey = "supports_client_text"
	ProductDeleteIDKey           = "product_to_delete"
	ProductDeleteTextKey         = "delete_text"
	ProductDeleteSuccessTextKey  = "delete_success_text"
	ProductDeleteURLKey          = "delete_url"
	ProductCategoriesKey         = "categories"
)

/// GenerateProductContent fills a string nested map with all product details and assets info
func (c *ContentController) generateProductContent(productData *restbackend.Product) map[string]interface{} {
	content := productData.Details
	for k, v := range productData.Assets {
		content[k] = v
	}

	content[ProductDetailPageKey] = fmt.Sprintf("/user-main/my-products/details?item-id=%s", productData.ID)
	content[ProductEditPageKey] = fmt.Sprintf("/user-main/my-products/edit?item-id=%s", productData.ID)
	content[NewProject] = fmt.Sprintf("/user-main/my-products/project-wizard?item-id=%s", productData.ID)
	content[ProductDeleteIDKey] = productData.ID
	content[ProductDeleteTextKey] = "It will delete all projects started from this product as well"
	content[ProductDeleteSuccessTextKey] = "Your product has been deleted"
	content[ProductDeleteURLKey] = "/user-main/my-products/delete"
	return content
}

// GetProductContent returns the selected product details and assets info.
func (c *ContentController) GetProductContent(productID string) (map[string]interface{}, error) {
	product, err := c.RESTBackend.GetProduct(productID)
	if err != nil {
		return nil, err
	}
	content := c.generateProductContent(product)
	return content, nil
}

// GetUserProductContent gathers the content of each product belonging to the specified user.
func (c *ContentController) GetUserProductContent(userID string) ([]map[string]interface{}, error) {
	products, err := c.RESTBackend.GetProductsByUserID(userID)
	if err != nil {
		return nil, err
	}

	productContent := make([]map[string]interface{}, len(products))
	for i, product := range products {
		productContent[i] = c.generateProductContent(product)
		productContent[i][ProductOwnerNameKey] = c.User.Settings[UserNameKey]
	}

	return productContent, nil
}

// GetProductsByCategory organizes product contents by categories. This is just a placeholder solution until proper search is introduced.
func (c *ContentController) GetProductsByCategory(userID string) (map[string]interface{}, error) {
	products, err := c.RESTBackend.GetProductsByUserID(userID)
	if err != nil {
		return nil, err
	}

	categorizedProducts := make(map[string]interface{})
	categories, err := c.RESTBackend.GetCategoriesMap()
	if err != nil {
		return nil, err
	}
	for k, v := range categories {
		if categorizedProducts[k] == nil {
			categorizedProducts[k] = make(map[string]interface{})
		}
		categorizedProducts[k].(map[string]interface{})["name"] = v
		for _, product := range products {
			for _, categoryKey := range product.Details[ProductCategoriesKey].([]interface{}) {
				if categoryKey.(string) == k {
					_, ok := categorizedProducts[k].(map[string]interface{})["products"]
					if !ok {
						categorizedProducts[k].(map[string]interface{})["products"] = make([]map[string]interface{}, 0)
					}
					categorizedProducts[k].(map[string]interface{})["products"] = append(categorizedProducts[k].(map[string]interface{})["products"].([]map[string]interface{}), c.generateProductContent(product))
					break
				}
			}
		}
	}
	return categorizedProducts, nil
}

// GetRecentProductsContent gathers the content of the latest 4 products
func (c *ContentController) GetRecentProductsContent(userID string) ([]map[string]interface{}, error) {
	products, err := c.RESTBackend.GetProductsByUserID(userID)
	if err != nil {
		return nil, err
	}

	limit := 4
	if len(products) < limit {
		limit = len(products)
	}
	productContent := make([]map[string]interface{}, limit)
	for i, product := range products {
		if limit == 0 {
			break
		}
		limit--
		productContent[i] = c.generateProductContent(product)
		productContent[i][ProductOwnerNameKey] = c.User.Settings[UserNameKey]
		productContent[i][ProductOwnerPageNameKey] = fmt.Sprintf("/user-main/profile?item-id=%s", c.User.ID)
		productContent[i][ProductDetailPageKey] = fmt.Sprintf("/user-main/product?item-id=%s", product.ID)
	}

	return productContent, nil
}

func (c *ContentController) GetRecommendedProductsContent(userID string) ([]map[string]interface{}, error) {
	products, err := c.RESTBackend.GetProductsByUserID(userID)
	if err != nil {
		return nil, err
	}

	limit := 4
	if len(products) < limit {
		limit = len(products)
	}
	productContent := make([]map[string]interface{}, limit)
	for i, product := range products {
		if limit == 0 {
			break
		}
		limit--
		productContent[i] = c.generateProductContent(product)
		productContent[i][ProductOwnerNameKey] = c.User.Settings[UserNameKey]
		productContent[i][ProductOwnerPageNameKey] = fmt.Sprintf("/user-main/profile?user=%s", c.User.ID)
		productContent[i][ProductDetailPageKey] = fmt.Sprintf("/user-main/product?item-id=%s", product.ID)
	}

	return productContent, nil
}
