"""
This module contains the Footer,
the page object for the page footer.
"""

from pages.page_object import PageObject
from selenium.webdriver.common.by import By


class Footer(PageObject):
    COPYRIGHT = (
      By.XPATH,
      "//a[@href='https://github.com/artofimagination/']"
    )

    FAQ = (
      By.XPATH,
      "//a[@class='nav-link' and @href='/resources/faq']"
    )

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def elementsPresent(self):
        result = dict()
        result["COPYRIGHT"] = self.findElement(self.COPYRIGHT)
        result["FAQ"] = self.findElement(self.FAQ)
        return result

    def clickFAQ(self):
        return self.clickButton(self.FAQ)

    def clickCopyright(self):
        copyrightButton = self.browser.find_element(*self.COPYRIGHT)
        copyrightButton.click()
        return self.browser.current_url
