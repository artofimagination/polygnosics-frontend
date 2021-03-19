"""
This module contains the UserMainPage,
the page object for the main user page.
"""

from pages.side_bar import SideBarPage
from pages.page_object import PageObject


class UserMainPage(PageObject):
    URL = "http://0.0.0.0:8085/user-main"

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)
        self.sidebar = SideBarPage(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        self.sidebar.elementsPresent()

    def title(self):
        return self.browser.title
