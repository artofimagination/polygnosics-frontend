"""
This module contains the NewsPage,
the page object for news page.
"""

from pages.page_object import PageObject
from pages.main_header import MainHeader
from pages.side_bar import Sidebar
from pages.news_feed import NewsFeed
from pages.content_header import ContentHeader
from pages.footer import Footer
import serverInfo


class NewsPage(PageObject):
    URL = "http://0.0.0.0:" + serverInfo.getPort() + "/resources/news"

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)
        self.sidebar = Sidebar(browser, pageObjects)
        self.mainHeader = MainHeader(browser, pageObjects)
        self.newsFeed = NewsFeed(browser, pageObjects)
        self.contentHeader = ContentHeader(browser, pageObjects)
        self.footer = Footer(browser, pageObjects)

    def elementsPresent(self, username):
        result = dict()
        result["main_header"] = self.mainHeader.elementsPresent()
        result["side_bar"] = self.sidebar.elementsPresent(username)
        result["news_feed"] = self.newsFeed.elementsPresent()
        result["content_header"] = \
            self.contentHeader.elementsPresent(
                "News", "Resources", "News")
        result["footer"] = self.footer.elementsPresent()
        return result

    def setPageObjects(self, pageObjects):
        super(NewsPage, self).setPageObjects(pageObjects)
        self.sidebar.pageObjects = pageObjects
        self.mainHeader.pageObjects = pageObjects
        self.newsFeed.pageObjects = pageObjects
        self.contentHeader.pageObjects = pageObjects
        self.footer.pageObjects = pageObjects
