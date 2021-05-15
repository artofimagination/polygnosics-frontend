"""
This module contains the TutorialsPage,
the page object for the tutorials page.
"""

from pages.main_header import MainHeader
from pages.side_bar import Sidebar
from pages.content_header import ContentHeader
from pages.page_object import PageObject
from selenium.webdriver.common.by import By


class TutorialsPage(PageObject):
    URL = "http://0.0.0.0:8085/resources/tutorials"

    VIDEO_ITEM_LINK = (
      By.XPATH,
      "//h4[contains(., 'Tutorials')]"
    )

    VIDEO_ITEM_TITLE = (
      By.XPATH,
      "//p[contains(., 'All tutorial videos and \
articles that will help you create your own content')]"
    )

    VIDEO_ITEM_SHORT = (
      By.XPATH,
      "//a[@href='/resources/tutorials' and contains(., 'More')]"
    )

    VIDEO_ITEM_EDIT = (
      By.XPATH,
      "//h4[contains(., 'FAQ')]"
    )

    VIDEO_ITEM_UPDATE_DATE = (
      By.XPATH,
      "//p[contains(., 'If you need a \
quick guide Frequently Asked Questions is where you want to start.')]"
    )

    ARTICLE_IMAGE = (
      By.XPATH,
      "//h4[contains(., 'Tutorials')]"
    )

    ARTICLE_TITLE = (
      By.XPATH,
      "//p[contains(., 'All tutorial videos and \
articles that will help you create your own content')]"
    )

    ARTICLE_SHORT = (
      By.XPATH,
      "//a[@href='/resources/tutorials' and contains(., 'More')]"
    )

    ARTICLE_EDIT = (
      By.XPATH,
      "//h4[contains(., 'FAQ')]"
    )

    ARTICLE_UPDATE_DATE = (
      By.XPATH,
      "//p[contains(., 'If you need a \
quick guide Frequently Asked Questions is where you want to start.')]"
    )

    ARTICLE_MORE = (
      By.XPATH,
      "//p[contains(., 'If you need a \
quick guide Frequently Asked Questions is where you want to start.')]"
    )

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)
        self.sidebar = Sidebar(browser, pageObjects)
        self.mainHeader = MainHeader(browser, pageObjects)
        self.contentHeader = ContentHeader(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self, userGroup):
        self.mainHeader.elementsPresent(userGroup)
        self.sidebar.elementsPresent(userGroup)
        self.contentHeader.elementsPresent(
            userGroup, "Tutorials", "Resources", "Tutorials")

        self.browser.find_element(*self.VIDEO_ITEM_LINK)
        self.browser.find_element(*self.VIDEO_ITEM_TITLE)
        self.browser.find_element(*self.VIDEO_ITEM_SHORT)
        self.browser.find_element(*self.VIDEO_ITEM_EDIT)
        self.browser.find_element(*self.VIDEO_ITEM_UPDATE_DATE)
        self.browser.find_element(*self.ARTICLE_LINK)
        self.browser.find_element(*self.ARTICLE_TITLE)
        self.browser.find_element(*self.ARTICLE_SHORT)
        self.browser.find_element(*self.ARTICLE_EDIT)
        self.browser.find_element(*self.ARTICLE_UPDATE_DATE)
        self.browser.find_element(*self.ARTICLE_MORE)

    def clickArticleMore(self):
        more = self.browser.find_element(*self.ARTICLE_MORE)
        more.click()
        return self.getPage()

    def clickNewTutorial(self):
        return self.contentHeader.clickNew()

    def clickEdit(self):
        edit = self.browser.find_element(*self.ARTICLE_EDIT)
        edit.click()
        return self.getPage()

    def title(self):
        return self.browser.title
