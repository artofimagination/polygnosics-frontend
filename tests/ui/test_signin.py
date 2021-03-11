import pytest
from pages.signin import SigninPage
from pages.user_main import UserMainPage

dataColumns = ("data", "expected")
createTestData = [
    ({
        'title': 'Polygnosics - Sign in'
    },
        "test-image:latest"),
]

ids = ['Page Loaded']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadingSigninPage(browser, data, expected):
    signinPage = SigninPage(browser)
    signinPage.load()

    assert data["title"] == signinPage.title()

    try:
        signinPage.elementsPresent()
    except Exception as e:
        assert f"{e}" == ""


dataColumns = ("data", "expected")
createTestData = [
    ({
        "email": "signup@test.com",
        "password": "asd123ASD"
    },
        "test-image:latest"),
]

ids = ['Success']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_Signin(browser, data, expected):
    signinPage = SigninPage(browser)
    signinPage.load()

    try:
        url = signinPage.signin(
          data["email"],
          data["password"])
        userMain = UserMainPage(browser)
        assert url == userMain.URL
    except Exception as e:
        assert f"{e}" == ""
