import pytest
import common
import traceback
from pages.signin import SigninPage
from pages.user_main import UserMainPage
from pages.error import ErrorPage


dataColumns = ("data", "expected")
createTestData = [
    (
        # Input data
        "",
        # Expected
        {
            "title": "Polygnosics - Sign in",
            "error": ""
        })
]

ids = ['Page Loaded']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadingSigninPage(browser, data, expected):
    try:
        signinPage = common.createPage(SigninPage, browser)
        signinPage.load()
        assert signinPage.title() == expected["title"]
        signinPage.elementsPresent()
    except Exception as e:
        traceback.print_exc()
        assert f"{e}" == expected["error"]


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
          "alertText": "",
          "error": "Message: Unable to locate element: .sweet-alert",
          "destinationPage": UserMainPage
      }),
    (
      # Input data
      {
          "email": "failure@test.com",
          "password": "123"
      },
      # Expected
      {
          "alertText": "Failed to login. Incorrect email or password",
          "error": "",
          "destinationPage": SigninPage
      }),
]

ids = ['Success', 'Failure']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_Signin(browser, data, expected):
    try:
        signinPage = common.createPage(SigninPage, browser)
        signinPage.load()
        page = signinPage.signin(
          data["email"],
          data["password"])

        try:
            (page, message) = signinPage.checkFailedSweetAlert()
            assert message == expected["alertText"]
        except Exception as e:
            assert str(e).strip() == expected["error"]

        if isinstance(page, ErrorPage):
            assert page.errorMessage() == expected["error"]
        assert type(page) == expected["destinationPage"]
    except Exception as e:
        traceback.print_exc()
        assert f"{e}" == expected["error"]
