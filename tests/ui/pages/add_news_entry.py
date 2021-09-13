"""
This module contains the AddNewsEntry,
the page object for adding news feed entry.
"""

from pages.main_header import MainHeader
from pages.page_object import PageObject
from pages.content_header import ContentHeader
from pages.footer import Footer
from pages.side_bar import Sidebar
from selenium.webdriver.common.by import By
import serverInfo


class AddNewsEntry(PageObject):
    URL = "http://0.0.0.0:" + serverInfo.getPort() +\
        "/resources/create-news-item"

    EDITOR_TOOLBAR = (
      By.XPATH,
      "//ul[@class='wysihtml-toolbar']"
    )

    EDIT_FIELD = (
      By.XPATH,
      "//textarea[@name='news']"
    )

    UPDATE_BUTTON = (
      By.XPATH,
      "//button[@id='create-new-item']"
    )

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)
        self.sidebar = Sidebar(browser, pageObjects)
        self.contentHeader = ContentHeader(browser, pageObjects)
        self.footer = Footer(browser, pageObjects)
        self.mainHeader = MainHeader(browser, pageObjects)

    def elementsPresent(self, userGroup):
        self.mainHeader.elementsPresent(userGroup)
        self.sidebar.elementsPresent(userGroup)
        self.contentHeader.elementsPresent(
            userGroup, "Create News", "Resources", "Create News")
        self.footer.elementsPresent()
        self.browser.find_element(*self.EDITOR_TOOLBAR)
        self.browser.find_element(*self.EDIT_FIELD)
        self.browser.find_element(*self.UPDATE_BUTTON)

    def updateText(self, text):
        edit = self.browser.find_element(*self.EDIT_FIELD)
        edit.send_keys(text)
        update = self.browser.find_element(*self.UPDATE_BUTTON)
        update.click()

        return self.getPage()
