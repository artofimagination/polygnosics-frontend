"""
This module contains the SideBarPage,
the page object for side bar element.
"""

from selenium.webdriver.common.by import By
from pages.page_object import PageObject


class SideBarPage(PageObject):
    URL = "http://0.0.0.0:8085/user-main"

    DASHBOARD_TEXT = (
        By.XPATH,
        "//li[@class='header' \
        and contains(text(), 'Dashboard')]")
    MAIN_DASH = (
        By.XPATH,
        "//span[contains(text(), 'Main Dashboards')]")
    USER_STATS = (
        By.XPATH,
        "//a[@href='/user-main/user-stats' \
        and contains(text(), 'User Stats')]")
    OVERALL_ITEM_STATS = (
        By.XPATH,
        "//a[@href='/user-main/products-projects' \
        and contains(text(), 'Overall Item Stats')]")
    ACCOUNTING_STATS = (
        By.XPATH,
        "//a[@href='/user-main/accounting' \
        and contains(text(), 'Accounting')]")
    UI_STATS = (
        By.XPATH,
        "//a[@href='/user-main/ui-stats' \
        and contains(text(), 'UI Stats')]")
    SYSTEM_HEALTH_STATS = (
        By.XPATH,
        "//a[@href='/user-main/system-health' \
        and contains(text(), 'System Health')]")
    MISUSE_STATS = (
        By.XPATH,
        "//a[@href='/user-main/misuse-metrics' \
        and contains(text(), 'Misuse Metrics')]")
    PROJECT_STATS = (
        By.XPATH,
        "//span[contains(text(), 'Project statistics')]")
    PRODUCT_STATS = (
        By.XPATH,
        "//span[contains(text(), 'Product statistics')]")
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

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        self.browser.find_element(*self.DASHBOARD_TEXT)
        self.browser.find_element(*self.MAIN_DASH)
        self.browser.find_element(*self.PROJECT_STATS)
        self.browser.find_element(*self.PRODUCT_STATS)

        self.browser.find_element(*self.PROJECT_TEXT)
        self.browser.find_element(*self.PROJECT_BROWSER)
        self.browser.find_element(*self.MY_PROJECTS)

        self.browser.find_element(*self.PRODUCT_TEXT)
        self.browser.find_element(*self.MARKETPLACE)
        self.browser.find_element(*self.MY_PRODUCTS)

        self.browser.find_element(*self.MAIN_DASH).click()
        self.browser.find_element(*self.USER_STATS)
        self.browser.find_element(*self.OVERALL_ITEM_STATS)
        self.browser.find_element(*self.ACCOUNTING_STATS)
        self.browser.find_element(*self.UI_STATS)
        self.browser.find_element(*self.SYSTEM_HEALTH_STATS)
        self.browser.find_element(*self.MISUSE_STATS)
