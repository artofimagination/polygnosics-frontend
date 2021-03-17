"""
This module contains IndexPage,
the page object for the main index page.
"""

from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait
from pages.page_object import PageObject
from pages.index_header import IndexPageHeader


class IndexPage(PageObject):
    URL = "http://0.0.0.0:8081/index"

    PRELOADER = (By.CLASS_NAME, 'loader-wrap')

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)
        self.header = IndexPageHeader(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def maximizeWindow(self):
        self.browser.maximize_window()

    def title(self):
        return self.browser.title

    def signUp(self):
        return self.header.goToSignup()

    def signIn(self):
        return self.header.goToSignin()

    def elementsPresent(self):
        self.header.elementsPresent()

    def waitUntilPreloaderDisappears(self):
        # wait for loading element to appear
        # - required to prevent prematurely checking if element
        #   has disappeared, before it has had a chance to appear
        WebDriverWait(self.browser, 0.1).until(
          EC.presence_of_element_located(self.PRELOADER))
        # then wait for the element to disappear
        WebDriverWait(self.browser, 5).until_not(
          lambda x: x.find_element(*self.PRELOADER).is_displayed())
