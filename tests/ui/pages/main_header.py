"""
This module contains the MainHeader,
the page object for the main header.
"""

from pages.page_object import PageObject
from selenium.webdriver.common.by import By


class MainHeader(PageObject):
    FULL_SCREEN = (
      By.XPATH,
      "//a[@title='Notifications']"
    )

    NOTIFICATIONS = (
      By.XPATH,
      "//a[@data-provide='fullscreen']"
    )

    UI_SETTNGS = (
      By.XPATH,
      "//a[@title='Settings']"
    )

    SEARCH_BAR = (
      By.XPATH,
      "//li[@class='search-bar']"
    )

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        result = dict()
        result["NOTIFICATIONS"] = self.findElement(self.NOTIFICATIONS)
        result["FULL_SCREEN"] = self.findElement(self.FULL_SCREEN)
        result["UI_SETTNGS"] = self.findElement(self.UI_SETTNGS)
        result["SEARCH_BAR"] = self.findElement(self.SEARCH_BAR)
        return result
