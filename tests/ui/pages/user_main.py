"""
This module contains the UserMainPage,
the page object for the main user page.
"""

from pages.main_header import MainHeader
from pages.side_bar import Sidebar
from pages.news_feed import NewsFeed
from pages.content_header import ContentHeader
from pages.footer import Footer
from pages.page_object import PageObject
from selenium.webdriver.common.by import By
import serverInfo


class UserMainPage(PageObject):
    URL = "http://0.0.0.0:" + serverInfo.getPort() + "/user-main"

    NEWS_TITLE = (
      By.XPATH,
      "//h4[@class='box-title' and contains(., 'News')]"
    )

    NEWS_TIMELINE_MORE = (
      By.XPATH,
      "//a[@href='/resources/news' and contains(., 'More')]"
    )

    RECENT_PRODUCTS_TITLE = (
      By.XPATH,
      "//h4[@class='box-title' and contains(., 'Recent products')]"
    )

    RECENT_PRODUCTS_MORE = (
      By.XPATH,
      "//a[@href='/user-main/store' and contains(., 'More')]"
    )

    RUNNING_PROJECTS_TITLE = (
      By.XPATH,
      "//h4[@class='box-title' and contains(., 'Running projects')]"
    )

    RUNNING_PROJECTS_MORE = (
      By.XPATH,
      "//a[@href='/user-main/my-projects' and contains(., 'More')]"
    )

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
        super(UserMainPage, self).setPageObjects(pageObjects)
        self.sidebar.pageObjects = pageObjects
        self.mainHeader.pageObjects = pageObjects
        self.newsFeed.pageObjects = pageObjects
        self.contentHeader.pageObjects = pageObjects
        self.footer.pageObjects = pageObjects

    def elementsPresent(self, username):
        result = dict()
        result["main_header"] = self.mainHeader.elementsPresent()
        result["side_bar"] = self.sidebar.elementsPresent(username)
        result["news_feed"] = self.newsFeed.elementsPresent()
        result["content_header"] = \
            self.contentHeader.elementsPresent(
                "Info board", "User", "Info board")
        result["footer"] = self.footer.elementsPresent()

        result["user_main"] = dict()
        result["user_main"]["NEWS_TITLE"] = self.findElement(self.NEWS_TITLE)
        result["user_main"]["NEWS_TIMELINE_MORE"] = \
            self.findElement(self.NEWS_TIMELINE_MORE)
        result["user_main"]["RECENT_PRODUCTS_TITLE"] = \
            self.findElement(self.RECENT_PRODUCTS_TITLE)
        result["user_main"]["RECENT_PRODUCTS_MORE"] = \
            self.findElement(self.RECENT_PRODUCTS_MORE)
        result["user_main"]["RUNNING_PROJECTS_TITLE"] = \
            self.findElement(self.RUNNING_PROJECTS_TITLE)
        result["user_main"]["RUNNING_PROJECTS_MORE"] = \
            self.findElement(self.RUNNING_PROJECTS_MORE)
        return result

    def clickNewsMore(self):
        return self.clickButton(self.NEWS_TIMELINE_MORE)

    def clickProductsMore(self):
        return self.clickButton(self.RECENT_PRODUCTS_MORE)

    def clickAddNewsEntry(self):
        return self.contentHeader.clickNew()

    def clickProjectsMore(self):
        return self.clickButton(self.RUNNING_PROJECTS_MORE)

    def title(self):
        return self.browser.title
