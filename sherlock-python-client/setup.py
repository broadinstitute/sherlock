# coding: utf-8

"""
    Sherlock

    The Data Science Platform's source-of-truth service. Note: this API will try to load and return associations in responses, so clients won't need to make as many requests. This behavior isn't recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).

    The version of the OpenAPI document: development
    Contact: dsp-devops@broadinstitute.org
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from setuptools import setup, find_packages  # noqa: H301

# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools
NAME = "sherlock-python-client"
VERSION = "v1.6.11"
PYTHON_REQUIRES = ">= 3.8"
REQUIRES = [
    "urllib3 >= 1.25.3, < 3.0.0",
    "python-dateutil >= 2.8.2",
    "pydantic >= 2",
    "typing-extensions >= 4.7.1",
]

setup(
    name=NAME,
    version=VERSION,
    description="Sherlock",
    author="DSP DevOps",
    author_email="dsp-devops@broadinstitute.org",
    url="git+https://github.com/broadinstitute/sherlock#subdirectory&#x3D;sherlock-python-client",
    keywords=["OpenAPI", "OpenAPI-Generator", "Sherlock"],
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    license="BSD-3-Clause",
    long_description_content_type='text/markdown',
    long_description="""\
    The Data Science Platform&#39;s source-of-truth service. Note: this API will try to load and return associations in responses, so clients won&#39;t need to make as many requests. This behavior isn&#39;t recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
    """,  # noqa: E501
    package_data={"sherlock_python_client": ["py.typed"]},
)
