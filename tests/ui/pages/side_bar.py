"""
This module contains the Sidebar,
the page object for side bar element.
"""

from selenium.webdriver.common.by import By
from pages.page_object import PageObject


class Sidebar(PageObject):
    # DASHBOARD_TEXT = (
    #     By.XPATH,
    #     "//li[@class='header' \
    #     and contains(text(), 'Dashboard')]")
    # MAIN_DASH = (
    #     By.XPATH,
    #     "//span[contains(text(), 'Main Dashboards')]")
    # USER_STATS = (
    #     By.XPATH,
    #     "//a[@href='/user-main/user-stats' \
    #     and contains(text(), 'User Stats')]")
    # OVERALL_ITEM_STATS = (
    #     By.XPATH,
    #     "//a[@href='/user-main/products-projects' \
    #     and contains(text(), 'Overall Item Stats')]")
    # ACCOUNTING_STATS = (
    #     By.XPATH,
    #     "//a[@href='/user-main/accounting' \
    #     and contains(text(), 'Accounting')]")
    # UI_STATS = (
    #     By.XPATH,
    #     "//a[@href='/user-main/ui-stats' \
    #     and contains(text(), 'UI Stats')]")
    # SYSTEM_HEALTH_STATS = (
    #     By.XPATH,
    #     "//a[@href='/user-main/system-health' \
    #     and contains(text(), 'System Health')]")
    # MISUSE_STATS = (
    #     By.XPATH,
    #     "//a[@href='/user-main/misuse-metrics' \
    #     and contains(text(), 'Misuse Metrics')]")
    # PROJECT_STATS = (
    #     By.XPATH,
    #     "//span[contains(text(), 'Project statistics')]")
    # PRODUCT_STATS = (
    #     By.XPATH,
    #     "//span[contains(text(), 'Product statistics')]")
    PROJECT_TEXT = (
        By.XPATH,
        "//li[@class='header' \
        and contains(text(), 'Projects')]")
    PROJECT_BROWSER = (
        By.XPATH,
        "//span[contains(text(), 'Browse projects')]")
    MY_PROJECTS = (
        By.XPATH,
        "//span[contains(text(), 'My Projects')]")
    PRODUCT_TEXT = (
        By.XPATH,
        "//li[@class='header' \
        and contains(text(), 'Products')]")
    MARKETPLACE = (
        By.XPATH,
        "//span[contains(text(), 'Marketplace')]")
    MY_PRODUCTS = (
        By.XPATH,
        "//span[contains(text(), 'My Products')]")
    RESOURCES_TEXT = (
        By.XPATH,
        "//li[@class='header' \
        and contains(text(), 'Resources')]")
    NEWS = (
        By.XPATH,
        "//span[contains(text(), 'News')]")
    DOCUMENTATION = (
        By.XPATH,
        "//span[contains(text(), 'Documentation')]")
    FILES = (
        By.XPATH,
        "//span[contains(text(), 'Files')]")
    ABOUT_TEXT = (
        By.XPATH,
        "//li[@class='header' \
        and contains(text(), 'About us')]")
    CONTACT = (
        By.XPATH,
        "//span[contains(text(), 'Contact')]")
    WHO_WE_ARE = (
        By.XPATH,
        "//span[contains(text(), 'Who we are?')]")

    HOME_BUTTON = (
        By.XPATH,
        "//a[@data-original-title='Home']")

    INBOX = (
        By.XPATH,
        "//a[@href='/user-main/mail-inbox' @data-original-title='Inbox']")

    SETTINGS = (
        By.XPATH,
        "//a[@href='/user-main/settings' @data-original-title='Settings']")

    LOGOUT = (
        By.XPATH,
        "//a[@href='/auth_logout' @data-original-title='Logout']")

    PROFILE_PHOTO = (
        By.XPATH,
        "//img[@src='/backend/avatar-test.jpg']")

    USERNAME = (
        By.XPATH,
        "//a[@class='px-20' and contains(text(), '{0}')]")

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def elementsPresent(self, username):
        result = dict()
        result["HOME_BUTTON"] = self.findElement(self.HOME_BUTTON)
        # self.browser.find_element(*self.DASHBOARD_TEXT)
        # self.browser.find_element(*self.MAIN_DASH)
        # self.browser.find_element(*self.PROJECT_STATS)
        # self.browser.find_element(*self.PRODUCT_STATS)

        result["PROJECT_TEXT"] = self.findElement(self.PROJECT_TEXT)
        result["PROJECT_BROWSER"] = self.findElement(self.PROJECT_BROWSER)
        result["MY_PROJECTS"] = self.findElement(self.MY_PROJECTS)

        result["PRODUCT_TEXT"] = self.findElement(self.PRODUCT_TEXT)
        result["MARKETPLACE"] = self.findElement(self.MARKETPLACE)
        result["MY_PRODUCTS"] = self.findElement(self.MY_PRODUCTS)

        # self.browser.find_element(*self.MAIN_DASH).click()
        # self.browser.find_element(*self.USER_STATS)
        # self.browser.find_element(*self.OVERALL_ITEM_STATS)
        # self.browser.find_element(*self.ACCOUNTING_STATS)
        # self.browser.find_element(*self.UI_STATS)
        # self.browser.find_element(*self.SYSTEM_HEALTH_STATS)
        # self.browser.find_element(*self.MISUSE_STATS)

        result["RESOURCES_TEXT"] = self.findElement(self.RESOURCES_TEXT)
        result["NEWS"] = self.findElement(self.NEWS)
        result["DOCUMENTATION"] = self.findElement(self.DOCUMENTATION)
        result["FILES"] = self.findElement(self.FILES)

        result["ABOUT_TEXT"] = self.findElement(self.ABOUT_TEXT)
        result["CONTACT"] = self.findElement(self.CONTACT)
        result["WHO_WE_ARE"] = self.findElement(self.WHO_WE_ARE)

        result["PROFILE_PHOTO"] = self.findElement(self.PROFILE_PHOTO)
        self.USERNAME = \
            (self.USERNAME[0], self.USERNAME[1].format(username))
        result["USERNAME"] = self.findElement(self.USERNAME)
        return result

    def clickProjectBrowser(self):
        return self.clickButton(self.PROJECT_BROWSER)

    def clickMyProjects(self):
        return self.clickButton(self.MY_PROJECTS)

    def clickMarketplace(self):
        return self.clickButton(self.MARKETPLACE)

    def clickMyProducts(self):
        return self.clickButton(self.MY_PRODUCTS)

    def clickNews(self):
        return self.clickButton(self.NEWS)

    def clickDocumentation(self):
        return self.clickButton(self.DOCUMENTATION)

    def clickFiles(self):
        return self.clickButton(self.FILES)

    def clickContact(self):
        return self.clickButton(self.CONTACT)

    def clickWhoWeAre(self):
        return self.clickButton(self.WHO_WE_ARE)

    def clickHome(self):
        return self.clickButton(self.HOME_BUTTON)
