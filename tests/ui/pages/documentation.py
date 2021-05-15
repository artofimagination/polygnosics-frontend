"""
This module contains the DocumentationPage,
the page object for the documentation page.
"""

from pages.main_header import MainHeader
from pages.side_bar import Sidebar
from pages.content_header import ContentHeader
from pages.page_object import PageObject
from selenium.webdriver.common.by import By


class DocumentationPage(PageObject):
    URL = "http://0.0.0.0:8085/resources/docs"

    TUTORIAL_TITLE = (
      By.XPATH,
      "//h4[contains(., 'Tutorials')]"
    )

    TUTORIAL_DESCRIPTION = (
      By.XPATH,
      "//p[contains(., 'All tutorial videos and \
articles that will help you create your own content')]"
    )

    TUTORIAL_MORE = (
      By.XPATH,
      "//a[@href='/resources/tutorials' and contains(., 'More')]"
    )

    FAQ_TITLE = (
      By.XPATH,
      "//h4[contains(., 'FAQ')]"
    )

    FAQ_DESCRIPTION = (
      By.XPATH,
      "//p[contains(., 'If you need a \
quick guide Frequently Asked Questions is where you want to start.')]"
    )

    FAQ_MORE = (
      By.XPATH,
      "//a[@href='/resources/faq' and contains(., 'More')]"
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
            userGroup, "Documentation", "Resources", "Documentation")

        self.browser.find_element(*self.TUTORIAL_TITLE)
        self.browser.find_element(*self.TUTORIAL_DESCRIPTION)
        self.browser.find_element(*self.TUTORIAL_MORE)
        self.browser.find_element(*self.FAQ_TITLE)
        self.browser.find_element(*self.FAQ_DESCRIPTION)
        self.browser.find_element(*self.FAQ_MORE)

    def clickTutorialsMore(self):
        more = self.browser.find_element(*self.TUTORIAL_MORE)
        more.click()
        return self.getPage()

    def clickFAQsMore(self):
        more = self.browser.find_element(*self.FAQ_MORE)
        more.click()
        return self.getPage()

    def title(self):
        return self.browser.title
