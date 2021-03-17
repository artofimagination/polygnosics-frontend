"""
This module contains the ErrorPage,
the page object for the errorpage.
"""

from selenium.webdriver.common.by import By


class ErrorPage:
    URL = "http://0.0.0.0:8081/error"

    ERROR_MESSAGE = (By.ID, 'error-message')

    def __init__(self, browser):
        self.browser = browser

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        self.sidebar.elementsPresent()

    def title(self):
        return self.browser.title

    def errorMessage(self):
        errorMessage = self.browser.find_element(*self.ERROR_MESSAGE)
        print(errorMessage.text)
        return errorMessage.get_attribute("text")
