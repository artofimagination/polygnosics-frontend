import pytest
import common
import traceback
from pages.signup import SignupPage
from pages.error import ErrorPage
from pages.index import IndexPage

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
            "alertText": "Registration successful",
            "error": "",
            "destinationPage": IndexPage
        }),
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
            "alertText": "User already exists",
            "error": "",
            "destinationPage": IndexPage
        }),
]

ids = ['Success', 'Failure']


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

        try:
            (page, message) = signupPage.checkSweetAlert()
            print(page)
            assert message == expected["alertText"]
        except Exception as e:
            assert str(e).strip() == expected["error"]

        if isinstance(page, ErrorPage):
            assert page.errorMessage() == expected["error"]
        assert type(page) == expected["destinationPage"]
    except Exception as e:
        traceback.print_exc()
        assert f"{e}" == expected["error"]
