import pytest
from pages.index_header import IndexPageHeader
from pages.signin import SigninPage
from pages.signup import SignupPage

dataColumns = ("data", "expected")
createTestData = [
    ({
        'title': 'Polygnosics',
    },
        "test-image:latest"),
]

ids = ['Elements exist']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_LoadingPageHeader(browser, data, expected):
    indexPageHeader = IndexPageHeader(browser)
    indexPageHeader.load()

    assert data["title"] == indexPageHeader.title()

    try:
        indexPageHeader.waitUntilPreloaderDisappears()
    except Exception as e:
        assert f"{e}" == ""

    try:
        url = indexPageHeader.goToSignup()
        signup = SignupPage(browser)
        assert url == signup.URL
    except Exception as e:
        assert f"{e}" == ""

    try:
        url = indexPageHeader.goToSignin()
        signin = SigninPage(browser)
        assert url == signin.URL
    except Exception as e:
        assert f"{e}" == ""
