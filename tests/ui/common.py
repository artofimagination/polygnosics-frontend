from pages.error import ErrorPage
from pages.user_main import UserMainPage
from pages.signup import SignupPage
from pages.signin import SigninPage
from pages.index import IndexPage


def createPageObjects(browser):
    pageObjects = dict()
    page = ErrorPage(browser)
    pageObjects[page.URL] = ErrorPage(browser)
    page = IndexPage(browser)
    pageObjects[page.URL] = IndexPage(browser)
    page = UserMainPage(browser)
    pageObjects[page.URL] = UserMainPage(browser)
    page = SigninPage(browser)
    pageObjects[page.URL] = SigninPage(browser)
    page = SignupPage(browser)
    pageObjects[page.URL] = SignupPage(browser)
    return pageObjects


def createPage(constructor, browser):
    page = constructor(browser, createPageObjects(browser))
    return page
