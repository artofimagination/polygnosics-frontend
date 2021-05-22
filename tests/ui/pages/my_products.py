"""
This module contains the MyProductsPage,
the page object for the my products page.
"""

from pages.main_header import MainHeader
from pages.side_bar import Sidebar
from pages.news_feed import NewsFeed
from pages.content_header import ContentHeader
from pages.footer import Footer
from pages.page_object import PageObject


class MyProductsPage(PageObject):
    URL = "http://0.0.0.0:8085/user-main/my-products"

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)
        self.sidebar = Sidebar(browser, pageObjects)
        self.mainHeader = MainHeader(browser, pageObjects)
        self.newsFeed = NewsFeed(browser, pageObjects)
        self.contentHeader = ContentHeader(browser, pageObjects)
        self.footer = Footer(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def setPageObjects(self, pageObjects):
        super(MyProductsPage, self).setPageObjects(pageObjects)
        self.sidebar.pageObjects = pageObjects
        self.mainHeader.pageObjects = pageObjects
        self.newsFeed.pageObjects = pageObjects
        self.contentHeader.pageObjects = pageObjects
        self.footer.pageObjects = pageObjects
