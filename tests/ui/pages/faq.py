"""
This module contains the FAQPage,
the page object for the FAQ page.
"""

from pages.main_header import MainHeader
from pages.side_bar import Sidebar
from pages.news_feed import NewsFeed
from pages.content_header import ContentHeader
from pages.footer import Footer
from pages.page_object import PageObject
import serverInfo


class FAQPage(PageObject):
    URL = "http://0.0.0.0:" + serverInfo.getPort() + "/resources/faq"

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
        super(FAQPage, self).setPageObjects(pageObjects)
        self.sidebar.pageObjects = pageObjects
        self.mainHeader.pageObjects = pageObjects
        self.newsFeed.pageObjects = pageObjects
        self.contentHeader.pageObjects = pageObjects
        self.footer.pageObjects = pageObjects
