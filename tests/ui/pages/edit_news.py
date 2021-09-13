"""
This module contains the EditNewsEntry,
the page object for editing news feed entry.
"""

from pages.main_header import MainHeader
from pages.page_object import PageObject
from pages.content_header import ContentHeader
from pages.footer import Footer
from pages.side_bar import Sidebar
from selenium.webdriver.common.by import By
import serverInfo


class EditNewsEntry(PageObject):
    URL = "http://localhost:" + serverInfo.getPort() +\
        "/resources/edit-news-item?id=d2b75a47-52f9-4401-af1b-a397005c14c3"

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

    CURRENT_FEED_TEXT = (
      By.XPATH,
      "//textarea[@name='news' and contains(., 'This is one news')]"
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
            userGroup, "Edit News", "Resources", "Edit News")
        self.footer.elementsPresent()
        self.browser.find_element(*self.EDITOR_TOOLBAR)
        self.browser.find_element(*self.EDIT_FIELD)
        self.browser.find_element(*self.UPDATE_BUTTON)
        self.browser.find_element(*self.CURRENT_FEED_TEXT)

    def updateText(self, text):
        edit = self.browser.find_element(*self.EDIT_FIELD)
        edit.send_keys(text)
        update = self.browser.find_element(*self.UPDATE_BUTTON)
        update.click()

        return self.getPage()
