"""
This module contains the NewsFeed,
the ppage object fro news feed element.
"""

from pages.page_object import PageObject
from selenium.webdriver.common.by import By


class NewsFeed(PageObject):
    NEWS_TIMELINE_2015 = (
      By.XPATH,
      "//span[@class='timeline__year' and contains(text(), '2015')]"
    )

    NEWS_TIMELINE_2015_ENTRY_DAY_1 = (
      By.XPATH,
      "//span[@class='timeline__day' and contains(text(), '5')]"
    )

    NEWS_TIMELINE_2015_ENTRY_MONTH_1 = (
      By.XPATH,
      "//span[@class='timeline__month' and contains(text(), 'Feb')]"
    )

    NEWS_TIMELINE_2015_ENTRY_1 = (
      By.XPATH,
      "//p[contains(text(), 'This is one news')]"
    )

    NEWS_TIMELINE_2015_ENTRY_EDIT_1 = (
      By.XPATH,
      "//a[contains(@href, \
'/resources/edit-news-item?id=d2b75a47-52f9-4401-af1b-a397005c14c3')]"
    )

    NEWS_TIMELINE_EDIT_ICON = (
      By.XPATH,
      "//i[@class='fa fa-pencil']"
    )

    NEWS_TIMELINE_2019 = (
      By.XPATH,
      "//span[@class='timeline__year' and contains(text(), '2019')]"
    )

    NEWS_TIMELINE_2019_ENTRY_DAY_1 = (
      By.XPATH,
      "//span[@class='timeline__day' and contains(text(), '6')]"
    )

    NEWS_TIMELINE_2019_ENTRY_MONTH_1 = (
      By.XPATH,
      "//span[@class='timeline__month' and contains(text(), 'Jun')]"
    )

    NEWS_TIMELINE_2019_ENTRY_1 = (
      By.XPATH,
      "//p[contains(text(), 'This is third news')]"
    )

    NEWS_TIMELINE_2019_ENTRY_EDIT_1 = (
      By.XPATH,
      "//a[contains(@href, \
'/resources/edit-news-item?id=d6d0381b-9c59-4cce-b9a0-15182a3ffef3')]"
    )

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def elementsPresent(self):
        result = dict()
        result["NEWS_TIMELINE_2015"] = \
            self.findElement(self.NEWS_TIMELINE_2015)
        result["NEWS_TIMELINE_2015_ENTRY_DAY_1"] = \
            self.findElement(self.NEWS_TIMELINE_2015_ENTRY_DAY_1)
        result["NEWS_TIMELINE_2015_ENTRY_MONTH_1"] = \
            self.findElement(self.NEWS_TIMELINE_2015_ENTRY_MONTH_1)
        result["NEWS_TIMELINE_2015_ENTRY_1"] = \
            self.findElement(self.NEWS_TIMELINE_2015_ENTRY_1)
        result["NEWS_TIMELINE_EDIT_ICON"] = \
            self.findChildElement(
                self.NEWS_TIMELINE_2015_ENTRY_EDIT_1,
                self.NEWS_TIMELINE_EDIT_ICON)
        result["NEWS_TIMELINE_2019"] = \
            self.findElement(self.NEWS_TIMELINE_2019)
        result["NEWS_TIMELINE_2019_ENTRY_DAY_1"] = \
            self.findElement(self.NEWS_TIMELINE_2019_ENTRY_DAY_1)
        result["NEWS_TIMELINE_2019_ENTRY_MONTH_1"] = \
            self.findElement(self.NEWS_TIMELINE_2019_ENTRY_MONTH_1)
        result["NEWS_TIMELINE_2019_ENTRY_1"] = \
            self.findElement(self.NEWS_TIMELINE_2019_ENTRY_1)
        result["NEWS_TIMELINE_EDIT_ICON"] = \
            self.findChildElement(
                self.NEWS_TIMELINE_2019_ENTRY_EDIT_1,
                self.NEWS_TIMELINE_EDIT_ICON)
        return result

    def editEntry(self):
        edit = self.browser.find_element(*self.NEWS_TIMELINE_2015_ENTRY_1)
        edit.click()

        return self.getPage()

    def checkFeedEntry(self, day, month, year, content):
        result = dict()
        timelineYear = (
            By.XPATH,
            f"//span[@class='timeline__year' and contains(text(), '{year}')]"
        )

        timelineDay = (
            By.XPATH,
            f"//span[@class='timeline__day' and contains(text(), '{day}')]"
        )
        result["TIMELINE_DAY"] = \
            self.findChildElement(timelineYear, timelineDay)

        timelineMonth = (
          By.XPATH,
          f"//span[@class='timeline__month' and contains(text(), '{month}')]"
        )
        result["TIMELINE_MONTH"] = \
            self.findChildElement(timelineYear, timelineMonth)

        timelineContent = (
          By.XPATH,
          f"//p[contains(text(), '{content}')]"
        )
        result["TIMELINE_CONTENT"] = \
            self.findElement(timelineYear, timelineContent)

        timelineEditLink = (
          By.XPATH,
          "//a[contains(@href, '/resources/edit-news-item?id=')]"
        )
        result["TIMELINE_EDIT_LINK"] = \
            self.findElement(timelineYear, timelineEditLink)
        result["TIMELINE_EDIT_ICON"] = \
            self.findElement(timelineYear, self.NEWS_TIMELINE_EDIT_ICON)
        return result
