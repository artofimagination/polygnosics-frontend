"""
This module contains the ContentHeader,
the page object for the content header.
"""

from pages.page_object import PageObject
from selenium.webdriver.common.by import By


class ContentHeader(PageObject):
    TITLE = (
      By.XPATH,
      "//h3[@class='page-title d-inline-block' and contains(., '{}')]"
    )

    HOME_LINK_PARENT = (
      By.XPATH,
      "//ol[@class='breadcrumb']"
    )

    HOME_ICON = (
      By.XPATH,
      "//i[@class='mdi mdi-home-outline']"
    )

    HOME_LINK = (
      By.XPATH,
      "//a[@href='/user-main']"
    )

    PARENT_PAGE_LINK = (
      By.XPATH,
      "//li[contains(., '{}')]"
    )

    PAGE_NAME = (
      By.XPATH,
      "//li[contains(., '{}')]"
    )

    NEW_BUTTON = (
      By.XPATH,
      "//a[@href='/resources/create-news-item']"
    )

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def elementsPresent(self, title, parent, pageName):
        result = dict()
        self.TITLE = \
            (self.TITLE[0], self.TITLE[1].format(title))
        result["TITLE"] = self.findElement(self.TITLE)
        result["HOME_ICON"] = \
            self.findChildElement(self.HOME_LINK_PARENT, self.HOME_ICON)
        result["HOME_LINK"] = \
            self.findChildElement(self.HOME_LINK_PARENT, self.HOME_LINK)
        self.PARENT_PAGE_LINK = \
            (self.PARENT_PAGE_LINK[0], self.PARENT_PAGE_LINK[1].format(parent))
        result["PARENT_PAGE_LINK"] = \
            self.findChildElement(self.HOME_LINK_PARENT, self.PARENT_PAGE_LINK)
        self.PAGE_NAME = \
            (self.PAGE_NAME[0], self.PAGE_NAME[1].format(pageName))
        result["PAGE_NAME"] = \
            self.findChildElement(self.HOME_LINK_PARENT, self.PAGE_NAME)
        result["NEW_BUTTON"] = self.findElement(self.NEW_BUTTON)
        return result

    def clickNew(self):
        new = self.browser.find_element(*self.NEW_BUTTON)
        new.click()

        return self.getPage()
