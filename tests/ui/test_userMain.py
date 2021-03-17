import pytest
import common
from pages.user_main import UserMainPage
from pages.signin import SigninPage
from pages.error import ErrorPage

dataColumns = ("data", "expected")
createTestData = [
    (
        # Input data
        {
          "email": "root@test.com",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User'
        }),
]

ids = ['Page Loaded']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadPage(browser, data, expected):
    signinPage = common.createPage(SigninPage, browser)
    signinPage.load()
    mainPage = signinPage.signin(
        data["email"],
        data["password"]
    )

    if isinstance(mainPage, ErrorPage):
        assert mainPage.errorMessage() == expected["error"]
    assert type(mainPage) == UserMainPage
    mainPage.load()

    try:
        mainPage.elementsPresent()
    except Exception as e:
        assert f"{e}" == ""
