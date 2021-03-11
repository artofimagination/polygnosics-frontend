"""
This module contains the UserMainPage,
the page object for the main user page.
"""


class UserMainPage:
    URL = "http://0.0.0.0:8081/user-main"

    def __init__(self, browser):
        self.browser = browser

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        pass
