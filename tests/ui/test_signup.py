import pytest
import common
import traceback
from pages.signup import SignupPage
from pages.error import ErrorPage

dataColumns = ("data", "expected")
createTestData = [
    (
        # Input data
        "",
        # Expected
        {
            "title": "Polygnosics - Sign up",
            "error": ""
        }),
]

ids = ['Page Loaded']


@pytest.mark.run(order=1)
@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadingSignupPage(browser, data, expected):
    signupPage = common.createPage(SignupPage, browser)
    signupPage.load()

    assert signupPage.title() == expected["title"]

    try:
        signupPage.elementsPresent()
    except Exception as e:
        traceback.print_exc()
        assert f"{e}" == expected["error"]


dataColumns = ("data", "expected")
createTestData = [
    (
        # Input data
        {
            "username": "signupUser",
            "email": "signup@test.com",
            "password": "asd123ASD",
            "password-repeat": "asd123ASD"
        },
        # Expected
        {
            "error": ""
        }),
]

ids = ['Success']


@pytest.mark.run(order=2)
@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_Signup(browser, data, expected):
    signupPage = common.createPage(SignupPage, browser)
    signupPage.load()

    try:
        page = signupPage.signup(
          data["username"],
          data["email"],
          data["password"],
          data["password-repeat"])
        if isinstance(page, ErrorPage):
            assert page.errorMessage() == expected["error"]
        assert type(page) == SignupPage
    except Exception as e:
        traceback.print_exc()
        assert f"{e}" == expected["error"]
