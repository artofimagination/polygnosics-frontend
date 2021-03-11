"""
This module contains SignupPage,
the page object for the signup page.
"""

from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait


class SignupPage:
    URL = "http://0.0.0.0:8081/auth_signup"

    USERNAME_FIELD = (By.ID, 'username')
    EMAIL_FIELD = (By.ID, 'email')
    PSW_FIELD = (By.ID, 'psw')
    PSW_REPEAT_FIELD = (By.ID, 'psw-repeat')
    GROUP_CHECKBOX = (By.ID, 'group')
    TC_CHECKBOX = (By.ID, 'terms')
    HEADER_1_TEXT = (
        By.XPATH,
        "//h2[text() ='Get started with Us']")
    HEADER_2_TEXT = (
        By.XPATH,
        "//p[text() ='Register a new membership']")
    GROUP_TEXT = (
        By.XPATH,
        "//label[text() ='I am a developer']")
    TC_TEXT = (
        By.XPATH,
        "//label[contains(text(), 'I agree to the')]")
    REGISTER_WITH_TEXT = (
        By.XPATH,
        "//p[text() ='- Register With -']")
    GOOGLE_BUTTON = (By.CLASS_NAME, 'fa-google-plus')
    GITHUB_BUTTON = (By.CLASS_NAME, 'fa-github')
    SIGNUP_BUTTON = (By.ID, 'register')

    def __init__(self, browser):
        self.browser = browser

    def load(self):
        self.browser.get(self.URL)

    def title(self):
        WebDriverWait(self.browser, 0.1).until(
          EC.presence_of_element_located(self.USERNAME_FIELD))
        return self.browser.title

    def elementsPresent(self):
        self.browser.find_element(*self.HEADER_1_TEXT)
        self.browser.find_element(*self.HEADER_2_TEXT)
        self.browser.find_element(*self.GROUP_TEXT)
        self.browser.find_element(*self.TC_TEXT)
        self.browser.find_element(*self.REGISTER_WITH_TEXT)

    def signup(self, username, email, password, passwordRepeat):
        usernameInput = self.browser.find_element(*self.USERNAME_FIELD)
        usernameInput.send_keys(username)

        emailInput = self.browser.find_element(*self.EMAIL_FIELD)
        emailInput.send_keys(email)

        passwordInput = self.browser.find_element(*self.PSW_FIELD)
        passwordInput.send_keys(password)

        passwordRepeatInput = self.browser.find_element(*self.PSW_REPEAT_FIELD)
        passwordRepeatInput.send_keys(passwordRepeat)

        signup = self.browser.find_element(*self.SIGNUP_BUTTON)
        signup.click()
