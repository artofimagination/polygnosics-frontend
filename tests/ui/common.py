from pages.error import ErrorPage
from pages.user_main import UserMainPage
from pages.signup import SignupPage
from pages.signin import SigninPage
from pages.index import IndexPage
from pages.add_news_entry import AddNewsEntry
from pages.news import NewsPage
from pages.project_browser import ProjectBrowserPage
from pages.my_projects import MyProjectsPage
from pages.marketplace import MarketplacePage
from pages.my_products import MyProductsPage
from pages.faq import FAQPage
from enum import Enum


class UserGroup(Enum):
    Root = "root"
    Admin = "admin"
    Developer = "developer"
    Client = "client"
    Visitor = "visitor"


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
    page = AddNewsEntry(browser)
    pageObjects[page.URL] = AddNewsEntry(browser)
    page = NewsPage(browser)
    pageObjects[page.URL] = NewsPage(browser)
    page = ProjectBrowserPage(browser)
    pageObjects[page.URL] = ProjectBrowserPage(browser)
    page = MyProjectsPage(browser)
    pageObjects[page.URL] = MyProjectsPage(browser)
    page = MarketplacePage(browser)
    pageObjects[page.URL] = MarketplacePage(browser)
    page = MyProductsPage(browser)
    pageObjects[page.URL] = MyProductsPage(browser)
    page = FAQPage(browser)
    pageObjects[page.URL] = FAQPage(browser)
    for pageObject in pageObjects.values():
        pageObject.setPageObjects(pageObjects)
    return pageObjects


def createPage(constructor, browser):
    page = constructor(browser, createPageObjects(browser))
    return page


def checkElementsResult(expected, result):
    resultString = ''
    if len(expected["elements"]) != len(result):
        raise Exception("Some elements are not checked")
    for groupKey, group in expected["elements"].items():
        if len(expected["elements"]) != len(result):
            raise Exception("Some elements are not checked")
        for key, value in group.items():
            if result[groupKey][key] != value:
                resultString += f"\n{groupKey} -> {key} failed. \
Result: {result[groupKey][key]} Expected: {value}"
    if resultString != '':
        raise Exception(resultString)


def testSidebar(page, expected):
    resultString = ''
    key = "project-browser"
    resultPage = page.sidebar.clickProjectBrowser()
    if resultPage.URL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultPage.URL} Expected: {expected[key]}"
    page.load()

    key = "my-projects"
    resultPage = page.sidebar.clickMyProjects()
    if resultPage.URL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultPage.URL} Expected: {expected[key]}"
    page.load()

    key = "marketplace"
    resultPage = page.sidebar.clickMarketplace()
    if resultPage.URL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultPage.URL} Expected: {expected[key]}"
    page.load()

    key = "my-products"
    resultPage = page.sidebar.clickMyProducts()
    if resultPage.URL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultPage.URL} Expected: {expected[key]}"
    page.load()

    key = "news"
    resultPage = page.sidebar.clickNews()
    if resultPage.URL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultPage.URL} Expected: {expected[key]}"
    page.load()
    if resultString != '':
        resultString = \
            "\n" + page.__class__.__name__ + \
            " Sidebar test failed" + resultString
        raise Exception(resultString)


def testFooter(page, expected):
    resultString = ''
    key = "copyright"
    resultURL = page.footer.clickCopyright()
    if resultURL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultURL} Expected: {expected[key]}"
    page.load()

    key = "faq"
    resultPage = page.footer.clickFAQ()
    if resultPage.URL != expected[key]:
        resultString += f"\n{key} failed. \
Result: {resultPage.URL} Expected: {expected[key]}"
    page.load()

    if resultString != '':
        resultString = \
            "\n" + page.__class__.__name__ + \
            " Sidebar test failed" + resultString
        raise Exception(resultString)
