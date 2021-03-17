"""
This module contains the SideBarPage,
the page object for side bar element.
"""

from selenium.webdriver.common.by import By
from pages.page_object import PageObject


class SideBarPage(PageObject):
    URL = "http://0.0.0.0:8081/user-main"

    PROJECT_STATS = (
        By.XPATH,
        "//a[@href='/user-main/project-stats' \
        and contains(text(), 'Project statistics')]")\


    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        pass
