import pytest
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
def test_LoadPage(browser, data, expected):
    page = UserMainPage(browser)
    page.load()

    assert data["title"] == page.title()

    try:
        page.elementsPresent()
    except Exception as e:
        assert f"{e}" == ""
