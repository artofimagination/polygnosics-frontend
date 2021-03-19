"""
This module contains IndexPageHeader,
the page object for the header of the main index page.
"""

from selenium.webdriver.common.by import By
from pages.page_object import PageObject


class IndexPageHeader(PageObject):
    URL = "http://0.0.0.0:8085/index"

    SIGNUP_BUTTON = (By.XPATH, "//a[@href='/auth_signup']")
    SIGNIN_BUTTON = (By.LINK_TEXT, 'Sign In')

    def __init__(self, browser, pageObjects):
        super().__init__(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def title(self):
        return self.browser.title

    def elementsPresent(self):
        pass

    def goToSignup(self):
        self.browser.execute_script(
            "window.scrollTo(0, document.body.scrollHeight*0.8);")
        signup = self.browser.find_element(*self.SIGNUP_BUTTON)
        signup.click()
        return self.getPage()

    def goToSignin(self):
        self.browser.execute_script(
            "window.scrollTo(0, document.body.scrollHeight*0.8);")
        signin = self.browser.find_element(*self.SIGNIN_BUTTON)
        signin.click()
        return self.getPage()
