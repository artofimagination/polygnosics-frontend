"""
This module contains the PageObject,
that is the base class for all page objects.
"""


class PageObject:
    def __init__(self, browser, pageObjects):
        self.browser = browser
        self.pageObjects = pageObjects

    def setPageObjects(self, pageObjects):
        self.pageObjects = pageObjects

    def getPage(self):
        if self.browser.current_url in self.pageObjects:
            return self.pageObjects[self.browser.current_url]
        return None

    def findElement(self, element):
        try:
            self.browser.find_element(*element)
        except Exception:
            return False
        return True

    def findChildElement(self, parent, element):
        try:
            parentElement = self.browser.find_element(*parent)
            parentElement.find_element(*element)
        except Exception:
            return False
        return True

    def clickButton(self, element):
        button = self.browser.find_element(*element)
        button.click()
        return self.getPage()
