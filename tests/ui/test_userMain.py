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
          "username": "root",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            "elements": {
                "user_main": {
                    "NEWS_TITLE": True,
                    "NEWS_TIMELINE_MORE": True,
                    "RECENT_PRODUCTS_TITLE": True,
                    "RECENT_PRODUCTS_MORE": True,
                    "RUNNING_PROJECTS_TITLE": True,
                    "RUNNING_PROJECTS_MORE": True
                },
                "main_header": {
                    "NOTIFICATIONS": True,
                    "FULL_SCREEN": True,
                    "UI_SETTNGS": True,
                    "SEARCH_BAR": True
                },
                "content_header": {
                    "TITLE": True,
                    "HOME_ICON": True,
                    "PARENT_PAGE_LINK": True,
                    "PAGE_NAME": True,
                    "NEW_BUTTON": True
                },
                "news_feed": {
                    "NEWS_TIMELINE_2015": True,
                    "NEWS_TIMELINE_2015_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": True,
                    "NEWS_TIMELINE_2019": True,
                    "NEWS_TIMELINE_2019_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": True
                },
                "side_bar": {
                    "PROJECT_TEXT": True,
                    "PROJECT_BROWSER": True,
                    "MY_PROJECTS": True,
                    "PRODUCT_TEXT": True,
                    "MARKETPLACE": True,
                    "MY_PRODUCTS": True,
                    "RESOURCES_TEXT": True,
                    "NEWS": True,
                    "DOCUMENTATION": True,
                    "FILES": True,
                    "ABOUT_TEXT": True,
                    "CONTACT": True,
                    "WHO_WE_ARE": True,
                    "PROFILE_PHOTO": True,
                    "USERNAME": True,
                    "HOME_BUTTON": True
                },
                "footer": {
                    "COPYRIGHT": True,
                    "FAQ": True
                }
            }
        }),
    (
        # Input data
        {
          "email": "admin@test.com",
          "username": "admin",
          "password": "123",
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            "elements": {
                "user_main": {
                    "NEWS_TITLE": True,
                    "NEWS_TIMELINE_MORE": True,
                    "RECENT_PRODUCTS_TITLE": True,
                    "RECENT_PRODUCTS_MORE": True,
                    "RUNNING_PROJECTS_TITLE": True,
                    "RUNNING_PROJECTS_MORE": True
                },
                "main_header": {
                    "NOTIFICATIONS": True,
                    "FULL_SCREEN": True,
                    "UI_SETTNGS": True,
                    "SEARCH_BAR": True
                },
                "content_header": {
                    "TITLE": True,
                    "HOME_ICON": True,
                    "PARENT_PAGE_LINK": True,
                    "PAGE_NAME": True,
                    "NEW_BUTTON": True
                },
                "news_feed": {
                    "NEWS_TIMELINE_2015": True,
                    "NEWS_TIMELINE_2015_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": True,
                    "NEWS_TIMELINE_2019": True,
                    "NEWS_TIMELINE_2019_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": True
                },
                "side_bar": {
                    "PROJECT_TEXT": True,
                    "PROJECT_BROWSER": True,
                    "MY_PROJECTS": True,
                    "PRODUCT_TEXT": True,
                    "MARKETPLACE": True,
                    "MY_PRODUCTS": True,
                    "RESOURCES_TEXT": True,
                    "NEWS": True,
                    "DOCUMENTATION": True,
                    "FILES": True,
                    "ABOUT_TEXT": True,
                    "CONTACT": True,
                    "WHO_WE_ARE": True,
                    "PROFILE_PHOTO": True,
                    "USERNAME": True,
                    "HOME_BUTTON": True
                },
                "footer": {
                    "COPYRIGHT": True,
                    "FAQ": True
                }
            }
        }),
    (
        # Input data
        {
          "email": "developer@test.com",
          "username": "developer",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            "elements": {
                "user_main": {
                    "NEWS_TITLE": True,
                    "NEWS_TIMELINE_MORE": True,
                    "RECENT_PRODUCTS_TITLE": True,
                    "RECENT_PRODUCTS_MORE": True,
                    "RUNNING_PROJECTS_TITLE": True,
                    "RUNNING_PROJECTS_MORE": True
                },
                "main_header": {
                    "NOTIFICATIONS": True,
                    "FULL_SCREEN": True,
                    "UI_SETTNGS": True,
                    "SEARCH_BAR": True
                },
                "content_header": {
                    "TITLE": True,
                    "HOME_ICON": True,
                    "PARENT_PAGE_LINK": True,
                    "PAGE_NAME": True,
                    "NEW_BUTTON": False
                },
                "news_feed": {
                    "NEWS_TIMELINE_2015": True,
                    "NEWS_TIMELINE_2015_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": False,
                    "NEWS_TIMELINE_2019": True,
                    "NEWS_TIMELINE_2019_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": False
                },
                "side_bar": {
                    "PROJECT_TEXT": True,
                    "PROJECT_BROWSER": True,
                    "MY_PROJECTS": True,
                    "PRODUCT_TEXT": True,
                    "MARKETPLACE": True,
                    "MY_PRODUCTS": True,
                    "RESOURCES_TEXT": True,
                    "NEWS": True,
                    "DOCUMENTATION": True,
                    "FILES": True,
                    "ABOUT_TEXT": True,
                    "CONTACT": True,
                    "WHO_WE_ARE": True,
                    "PROFILE_PHOTO": True,
                    "USERNAME": True,
                    "HOME_BUTTON": True
                },
                "footer": {
                    "COPYRIGHT": True,
                    "FAQ": True
                }
            }
        }),
    (
        # Input data
        {
          "email": "client@test.com",
          "username": "client",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            "elements": {
                "user_main": {
                    "NEWS_TITLE": True,
                    "NEWS_TIMELINE_MORE": True,
                    "RECENT_PRODUCTS_TITLE": True,
                    "RECENT_PRODUCTS_MORE": True,
                    "RUNNING_PROJECTS_TITLE": True,
                    "RUNNING_PROJECTS_MORE": True
                },
                "main_header": {
                    "NOTIFICATIONS": True,
                    "FULL_SCREEN": True,
                    "UI_SETTNGS": True,
                    "SEARCH_BAR": True
                },
                "content_header": {
                    "TITLE": True,
                    "HOME_ICON": True,
                    "PARENT_PAGE_LINK": True,
                    "PAGE_NAME": True,
                    "NEW_BUTTON": False
                },
                "news_feed": {
                    "NEWS_TIMELINE_2015": True,
                    "NEWS_TIMELINE_2015_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": False,
                    "NEWS_TIMELINE_2019": True,
                    "NEWS_TIMELINE_2019_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": False
                },
                "side_bar": {
                    "PROJECT_TEXT": True,
                    "PROJECT_BROWSER": True,
                    "MY_PROJECTS": True,
                    "PRODUCT_TEXT": True,
                    "MARKETPLACE": True,
                    "MY_PRODUCTS": True,
                    "RESOURCES_TEXT": True,
                    "NEWS": True,
                    "DOCUMENTATION": True,
                    "FILES": True,
                    "ABOUT_TEXT": True,
                    "CONTACT": True,
                    "WHO_WE_ARE": True,
                    "PROFILE_PHOTO": True,
                    "USERNAME": True,
                    "HOME_BUTTON": True
                },
                "footer": {
                    "COPYRIGHT": True,
                    "FAQ": True
                }
            }
        }
    ),
    (
        # Input data
        {
          "email": "visitor@test.com",
          "username": "visitor",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            "elements": {
                "user_main": {
                    "NEWS_TITLE": True,
                    "NEWS_TIMELINE_MORE": True,
                    "RECENT_PRODUCTS_TITLE": True,
                    "RECENT_PRODUCTS_MORE": True,
                    "RUNNING_PROJECTS_TITLE": True,
                    "RUNNING_PROJECTS_MORE": True
                },
                "main_header": {
                    "NOTIFICATIONS": True,
                    "FULL_SCREEN": True,
                    "UI_SETTNGS": True,
                    "SEARCH_BAR": True
                },
                "content_header": {
                    "TITLE": True,
                    "HOME_ICON": True,
                    "PARENT_PAGE_LINK": True,
                    "PAGE_NAME": True,
                    "NEW_BUTTON": False
                },
                "news_feed": {
                    "NEWS_TIMELINE_2015": True,
                    "NEWS_TIMELINE_2015_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2015_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": False,
                    "NEWS_TIMELINE_2019": True,
                    "NEWS_TIMELINE_2019_ENTRY_DAY_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_MONTH_1": True,
                    "NEWS_TIMELINE_2019_ENTRY_1": True,
                    "NEWS_TIMELINE_EDIT_ICON": False
                },
                "side_bar": {
                    "PROJECT_TEXT": True,
                    "PROJECT_BROWSER": True,
                    "MY_PROJECTS": True,
                    "PRODUCT_TEXT": True,
                    "MARKETPLACE": True,
                    "MY_PRODUCTS": True,
                    "RESOURCES_TEXT": True,
                    "NEWS": True,
                    "DOCUMENTATION": True,
                    "FILES": True,
                    "ABOUT_TEXT": True,
                    "CONTACT": True,
                    "WHO_WE_ARE": True,
                    "PROFILE_PHOTO": True,
                    "USERNAME": True,
                    "HOME_BUTTON": True
                },
                "footer": {
                    "COPYRIGHT": True,
                    "FAQ": True
                }
            }
        }
    ),
]

ids = [
  'Page Loaded Root',
  'Page Loaded Admin',
  'Page Loaded Developer',
  'Page Loaded Client',
  'Page Loaded Visitor']


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
        result = mainPage.elementsPresent(data["username"])
        common.checkElementsResult(expected, result)
    except Exception as e:
        pytest.fail(f"{e}")


dataColumns = ("data", "expected")

createTestData = [
    (
        # Input data
        {
          "email": "root@test.com",
          "username": "root",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            'url': "http://0.0.0.0:8085/resources/create-news-item"
        }
    ),
]

ids = ['Success']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_ClickAddNewsEntry(browser, data, expected):
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
    newEntryPage = mainPage.clickAddNewsEntry()
    assert newEntryPage.URL == expected['url']


dataColumns = ("data", "expected")

createTestData = [
    (
        # Input data
        {
          "email": "root@test.com",
          "username": "root",
          "password": "123"
        },
        # Expected
        {
            'title': 'Polygnosics - User',
            'url': "http://0.0.0.0:8085/resources/news"
        }
    ),
]

ids = ['Success']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_ClickNewsMore(browser, data, expected):
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
    newsPage = mainPage.clickNewsMore()
    assert newsPage.URL == expected['url']


dataColumns = ("data", "expected")

createTestData = [
    (
        # Input data
        {
          "email": "root@test.com",
          "username": "root",
          "password": "123"
        },
        # Expected
        {
            'project-browser': "http://0.0.0.0:8085/user-main/project-browser",
            'my-projects': "http://0.0.0.0:8085/user-main/my-projects",
            'marketplace': "http://0.0.0.0:8085/user-main/store",
            'my-products': "http://0.0.0.0:8085/user-main/my-products",
            'news': "http://0.0.0.0:8085/resources/news",
            'copyright': "https://github.com/artofimagination/",
            'faq': "http://0.0.0.0:8085/resources/faq"
        }
    ),
]

ids = ['Success']


@pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
def test_CommonComponents(browser, data, expected):
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
    common.testSidebar(mainPage, expected)
    common.testFooter(mainPage, expected)

# dataColumns = ("data", "expected")

# createTestData = [
#     (
#         # Input data
#         {
#           "email": "root@test.com",
#           "username": "root",
#           "password": "123"
#         },
#         # Expected
#         {
#             'title': 'Polygnosics - User',
#             'url': "http://0.0.0.0:8085/user-main/my-projects"
#         }
#     ),
# ]

# ids = ['Success']


# @pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
# def test_ClickProjectsMore(browser, data, expected):
#     signinPage = common.createPage(SigninPage, browser)
#     signinPage.load()
#     mainPage = signinPage.signin(
#         data["email"],
#         data["password"]
#     )

#     if isinstance(mainPage, ErrorPage):
#         assert mainPage.errorMessage() == expected["error"]
#     assert type(mainPage) == UserMainPage
#     mainPage.load()
#     projectsPage = mainPage.clickProjectsMore()
#     assert projectsPage.URL == expected['url']


# dataColumns = ("data", "expected")

# createTestData = [
#     (
#         # Input data
#         {
#           "email": "root@test.com",
#           "username": "root",
#           "password": "123"
#         },
#         # Expected
#         {
#             'title': 'Polygnosics - User',
#             'url': "http://0.0.0.0:8085/user-main/my-products"
#         }
#     ),
# ]

# ids = ['Success']


# @pytest.mark.parametrize(dataColumns, createTestData, ids=ids)
# def test_ClickProductsMore(browser, data, expected):
#     signinPage = common.createPage(SigninPage, browser)
#     signinPage.load()
#     mainPage = signinPage.signin(
#         data["email"],
#         data["password"]
#     )

#     if isinstance(mainPage, ErrorPage):
#         assert mainPage.errorMessage() == expected["error"]
#     assert type(mainPage) == UserMainPage
#     mainPage.load()
#     productsPage = mainPage.clickProductsMore()
#     assert productsPage.URL == expected['url']
