import pytest
from pages.signup import SignupPage

dataColumns = ("data", "expected")
createTestData = [
    ({
        'title': 'Polygnosics - Sign up'
    },
        "test-image:latest"),
]

ids = ['Page Loaded']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadingSignupPage(browser, data, expected):
    signupPage = SignupPage(browser)
    signupPage.load()

    assert data["title"] == signupPage.title()

    try:
        signupPage.elementsPresent()
    except Exception as e:
        assert f"{e}" == ""


dataColumns = ("data", "expected")
createTestData = [
    ({
        "username": "signupUser",
        "email": "signup@test.com",
        "password": "asd123ASD",
        "password-repeat": "asd123ASD"
    },
        "test-image:latest"),
]

ids = ['Success']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_Signup(browser, data, expected):
    signupPage = SignupPage(browser)
    signupPage.load()

    try:
        signupPage.signup(
          data["username"],
          data["email"],
          data["password"],
          data["password-repeat"])
    except Exception as e:
        assert f"{e}" == ""
