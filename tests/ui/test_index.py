import pytest
import common
import traceback
from pages.index import IndexPage
from pages.error import ErrorPage
from pages.signin import SigninPage
from pages.signup import SignupPage

dataColumns = ("data", "expected")
createTestData = [
    (
        # Input data
        "",
        # Expected
        {
            "title": "Polygnosics",
            "error": "",
        }),
]

ids = ['Elements exist']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadingPage(browser, data, expected):
    try:
        indexPage = common.createPage(IndexPage, browser)
        indexPage.load()
        assert indexPage.title() == expected["title"]
        indexPage.waitUntilPreloaderDisappears()

        page = indexPage.signUp()
        if isinstance(page, ErrorPage):
            assert page.errorMessage() == expected["error"]
        assert type(page) == SignupPage

        page = indexPage.signIn()
        if isinstance(page, ErrorPage):
            assert page.errorMessage() == expected["error"]
        assert type(page) == SigninPage
    except Exception as e:
        traceback.print_exc()
        assert f"{e}" == expected["error"]
