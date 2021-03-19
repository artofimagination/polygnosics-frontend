"""
This module contains SigninPage,
the page object for the signup page.
"""

from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from pages.page_object import PageObject


class SigninPage(PageObject):
    URL = "http://0.0.0.0:8085/auth_login"

    EMAIL_FIELD = (By.ID, 'email')
    PSW_FIELD = (By.ID, 'password')
    HEADER_1_TEXT = (
        By.XPATH,
        "//h2[text() ='Get started with Us']")
    HEADER_2_TEXT = (
        By.XPATH,
        "//p[text() ='Sign in to start your session']")
    REMEMBER_ME_TEXT = (
        By.XPATH,
        "//label[text() ='Remember Me']")
    FORGOT_PASSWORD_TEXT = (
        By.XPATH,
        "//a[contains(text(), ' Forgot password?')]")
    REGISTER_WITH_TEXT = (
        By.XPATH,
        "//p[text()='- Sign in With -']")
    NO_ACCOUNT_TEXT = (
        By.XPATH,
        "//p[contains(text(), 'have an account?')]")
    SIGNIN_BUTTON = (
        By.XPATH,
        "//button[text() ='SIGN IN']")
    FAILURE_SWEET_ALERT = (
        By.CLASS_NAME,
        "sweet-alert")
    FAILURE_SWEET_ALERT_MESSAGE = (
        By.XPATH,
        "//p[text() = 'Failed to login. Incorrect email or password']")
    FAILURE_SWEET_ALERT_OK = (
        By.XPATH,
        "//button[@class='confirm']")

    def __init__(self, browser, pageObjects=None):
        super().__init__(browser, pageObjects)

    def title(self):
        return self.browser.title

    def load(self):
        self.browser.get(self.URL)

    def elementsPresent(self):
        self.browser.find_element(*self.HEADER_1_TEXT)
        self.browser.find_element(*self.HEADER_2_TEXT)
        self.browser.find_element(*self.FORGOT_PASSWORD_TEXT)
        self.browser.find_element(*self.REGISTER_WITH_TEXT)
        self.browser.find_element(*self.REMEMBER_ME_TEXT)
        self.browser.find_element(*self.NO_ACCOUNT_TEXT)

    def checkFailedSweetAlert(self):
        self.browser.find_element(*self.FAILURE_SWEET_ALERT)
        message = \
            self.browser.find_element(*self.FAILURE_SWEET_ALERT_MESSAGE).text
        ok = self.browser.find_element(*self.FAILURE_SWEET_ALERT_OK)
        ActionChains(self.browser).move_to_element(ok).click(ok).perform()
        return (self.getPage(), message)

    def signin(self, email, password):
        emailInput = self.browser.find_element(*self.EMAIL_FIELD)
        emailInput.send_keys(email)

        passwordInput = self.browser.find_element(*self.PSW_FIELD)
        passwordInput.send_keys(password)

        signin = self.browser.find_element(*self.SIGNIN_BUTTON)
        signin.click()

        return self.getPage()
