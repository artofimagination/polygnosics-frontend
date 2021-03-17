"""
This module contains the PageObject,
that is the base class for all page objects.
"""


class PageObject:
    def __init__(self, browser, pageObjects):
        self.browser = browser
        self.pageObjects = pageObjects

    def getPage(self):
        if self.browser.current_url in self.pageObjects:
            return self.pageObjects[self.browser.current_url]
        return None
