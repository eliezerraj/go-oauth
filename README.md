# go-oauth

POC for test purposes.

## Endpoints

+ GET /header

+ GET /info

+ POST /oauth_credential
+ POST /oauth_credential_hs256

        {
            "user":"admin",
            "password":"admin"
        }

+ GET /tokenValidation/{token}

+ GET /tokenValidation_hs256/{token}

+ POST /refresh_token
+ POST /refresh_token_256

        {
            "token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl91c2UiOiJhY2Nlc3MiLCJpc3MiOiJnby1vYXV0aCIsInZlcnNpb24iOiIzIiwiand0X2lkIjoiOTNjYTRlMjItZTBjZC00NjEwLTkwYzgtOWRjZDI5MDlkZjdkIiwidXNlcm5hbWUiOiJhZG1pbiIsInNjb3BlIjpbImFkbWluIl0sImV4cCI6MTc0MTY5MTEzNH0.IkmQxyIOS00mAvQ96LdMgECTpkdowXalKt_8wnUlQqqZPDXWXh164AabtzVyX2sKSmq9MHKnbmz2S4YxpnZJ8hqcqNDlFS3FnZwSz4aD3_u1u8nfuesOodeqeoJ3YPGaYapZVUVScCJdp3kMI3e422pkgK6v3Vfq2GuEpjZeLsGsNY_XmqEI7CX8Ss5ojyDRr1RLRdNoKQWyugJ8d_PzPrd0G6MSwgzX9ypd6v_W_rujbVfUY677ru_NW3a-vTQB5puvLfYIOZrcdHOHu_cwz40Ry2adimN4BiRa1pgznTMTO-f-cqziMrnuhNnDt0t8mvAWKrYm0Emp3r5a12N5Hv0QR0DU4NnWDC_xcTcrekoVLnMGsILJ7UbjDna82-Wboy9ikl1Mo_aL-UIALD6XMi5NHl2dtlfNRdCTmRhLpbjr9JlUaFa1TjqjZM9ka57Qo13naIgfMI4rx9YTQGF32Z2O4YCAOVTiXilnUkATfanx9XAa2KLullsPK4lbh3aeNo7hcUdxCyxokUJlcrwzW6_a0-7V58A_RCyAKwjAthJI-TwRTXLFDOBdRS0u4r8YO_NZBo7nlqvaiJV6xP9PQ4vFzdnpwowVFEy2wYG63SiQ8KPFwf2iNHDrb8c5QUoDYqTckDXwu1EBKpiik5lxrIUN3vaB3-f4dvor3biHmBI"
        }

+ GET /wellKnown/1

+ POST /validationTokenSignedPubKey

        {
            "token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl91c2UiOiJhY2Nlc3MiLCJpc3MiOiJnby1vYXV0aCIsInZlcnNpb24iOiIzIiwiand0X2lkIjoiYjYyM2Q5M2YtZmJiYS00NDE4LThlMjAtOTY5NDA1ZTYxMGFhIiwidXNlcm5hbWUiOiJhZG1pbiIsInNjb3BlIjpbImFkbWluIl0sImV4cCI6MTc0MTY5MzYwMH0.C1q44Ryxy-PGAzwTa4LbbWYdJuxbwvTPDnpodZDKuBb5862l8l6L0mg14rPpOJN9I8AvLfUdOPSh9eH0QeR6De6AlWIdQTloapki43MP6hoFW_PFuxUG53kY2fOfJaotYbO7Ob0Wm3o_JQRVzwQAp4PeVvHPCxsI3zuWsPpbzCQuvxhMbxCHq3wX2NW19ppZF9kVqncGqRnpp7-4jqnztNAOluEJWMi4F9IUM1gEtjt9DIi7IKCz_I4auORNCY9YmJfFpbR_IYKkNHmvrcQvBBvN-IYIvXbPMx1NYJbpn6EsvFt-apJOWgdI8iz-BfTgoXPffKi1n_Mfy8XTsu3YrAZxqkxHHttIU4MO2UFMlpqUGmAQew8BqDLDqmv3bJ1PEqreTCGLf7H93MVmjsEUF9oKgSe4aKkeLpojT0qxZWdxo9hoQI8nC199qLvC7fwNtNooYPRaRy-l1ImAsy6P-eOuAe5Hn06s3vMOn_LZgX6c6eGxKioG0jyfPGQnt_ygGk5mhX4LLm4EjWkg7Sj_XBSfUCWwS0_DiHBj5qt-RO__7r8BQqVKTOX3RSEhYTHEnYiOzaR4gET1avE_FOe69OtnFxlbRTkSDBIB4tVcAGS4IRwgv359-I0FyGtY8UdDLhwxfvDQ_IPLQTRaijbfAaxDVGILtURtywCbHR7OEug",
            "rsa_public_key_b64":"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUFwSjAxd1AyZUE3a0tqUldIemV2dApjMWlab0R3Ym93cExkOXAzWGJDemJQbXZJYS9NK1VXbHBSZlJrTmZnaWt1YmUyV25EN0o1YjZmWmZiN2RYZUhKCjk1VHBlOG55V3FlWDFIKzlySXMyU3Iya1FyMVFzQTNYK3hDamZuaVkwK3pKbysyS2wxc3YyMU54QmdYOEhBVUkKSkxaanNSVmRuQUROOWlLVkR2NlVobkpONU45aHpPQlVISlVKc2ljWUlaQlFaaWRMSUhvSTlUcVVwaDk5emQwaApWOVlvYTNXazhRSDlJSVBLZG94dk45UklTZEpFWkh4UHMwcTFYSVZHY3Q5cWx3c0U5ZzdMdDh0eWlsYXY2RGFqCm9YUVZBR29LdU9uVERFNzBsYnRBQll1S0pLTzh5eDU2WUh4U21oMVFNbjg4c1p4dFRpbWZlWmJkd0U2Ykx5c24KQlM5NVNyNTMzeVVWT25iTFVFZnd4V3VrOWR1VWZXU3h5OUZMZlVwNnQ2M2liQmJlL0xrcVhRY0xsNExRZngydQp5NWR3YmF5THV0dnQyWWpGZmxQeVdOSkR2N2RJUWdCMFR0TE9ObkxWcFNlV1JEcEZ4bFlOVktHeWd0WXRtQTN6CnJ4VnIzckphZ2wyazRaMURCN3BjeUE0V3k3OEx2dlh6aHVyUk9CNkFRZzVDWmY1bkNwUFVlUC9YdXJyZkhOd2gKTVd3NDd2M0V2U3RVWVE1SkVKWnJQYWxoV0g5SksxUmVaVnlObmVCM0pyazhkWDI2SXBqaHFPUE5kMUFMSEdIeApqTGs3R1F4Y3p3NktHRVJmOEt2Tmx2b2JzZlg0Q1FlVGJ3QURhUWhBcGRPVkkxYVN0ay9KUkNvUER5SVM0dmU4CkRSa1ZONEJLaGprQ1NIdGR3NGxRS1k4Q0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo"
        }

+ POST /verifyCertCRL

        {
            "ca_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUdxekNDQkpPZ0F3SUJBZ0lVSGhwSXhFQ1pveHlFeFVuajZReVg5cUVaMWtZd0RRWUpLb1pJaHZjTkFRRUwKQlFBd2VqRXJNQ2tHQTFVRUF3d2laMjh0WVhWMGFEQXVZWEpqYUdsMFpXTjBkWEpsTG1OaGNtRmthSEpoY3k1cApiekVOTUFzR0ExVUVDZ3dFWkc5amF6RVZNQk1HQTFVRUN3d01ZWEpqYUdsMFpXTjBkWEpsTVFzd0NRWURWUVFHCkV3SkNVakVMTUFrR0ExVUVDQXdDVTFBeEN6QUpCZ05WQkFjTUFsTlFNQjRYRFRJME1USXhOVEl6TXpBME5Wb1gKRFRNME1USXhNekl6TXpBME5Wb3dhREVMTUFrR0ExVUVCaE1DUWxJeEN6QUpCZ05WQkFnTUFsTlFNUXN3Q1FZRApWUVFIREFKVFVERVhNQlVHQTFVRUNnd09ZMnhwWlc1MFpUQXhMbVJ2WTJzeERUQUxCZ05WQkFzTUJHRnlZMmd4CkZ6QVZCZ05WQkFNTURtTnNhV1Z1ZEdVd01TNWtiMk5yTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEEKTUlJQ0NnS0NBZ0VBcUlhUndYV3p0YlY5Mkkyd1gzUDc5MVdLRTVGZ3VoQStmYXBHNkhrSjVGWWRrWmRiQmNTVgpZeDVHVXBmNTR6ZmUvb3hiSG1jclY3OUlWU3M4SWNSTkMvRm1FaVk1c2syUm56TFN5am03Nmx6US9JeWJKZ2dnCkI2TzdPRVlqME5LUWFTajY4YW0zWUVSTjdHR1NKejFEYVhKdlA4eStod1BrWDRXZUdPVnVTNW03WElTWlJhdisKenlQcmxmeWw3R3ZiVVF0SHhPUE84Y0xNYnJpZXNaQy9qTE9oTDRHNzQxb1hJVlFKYVo0bGxzOUZ5eFhZUFFpQworRkl4UGxkN3BjUXhxa1JadWhaM2dHTjZtUGRZdGdkK2JQNFl2VEQwZEUvMXNSZGppSEowRkpabmpuWllPRnRhCnhsN0NMTzhQNWcrYTJ4WlFIbVJqTm5TMmhIb2IyU1JNcDFPV3NWbHRRY1Z4bTVaT2ttRVB1VW9tWUp5VzlpY2MKZGR3R2x4RllyaERQMHZYQlpVYlJsN1VMbmZYbkRCL1d5cldOMDl5OUpjK2g3WlhxSmwzMCtxUWl6NUxrRHdmbwpXWWtIRVdrMVgvcTE3bTl4Yk1qUlBuTm5lWFZGMnA1b203MmRnRFU4aE4yUTRzYnR2L0hxUkhabm1meU00cXdLCm9ZdnZPVWVheXNoVis2Q01yS3BESCtFVFZpU1pYRUVmMjlkL0hDckFXclUxcVdRQ0tsZlpqQndxR3dqSHB6MFgKQ1k0RE95RVBCeFd4RmNKbktSUkhJNzRlRWtma0xqRk5WWEVudnB6YytkV0FSUmYxbFE2MzF0bVpYT244aS9GRgp1bUIxQlBpeFJsbldQeHNKVis3VFI2dFBLMnFZWFFGWG1NaTF6T2k4OXR3dmNNSjlnbmVaZmhrQ0F3RUFBYU9DCkFUa3dnZ0UxTUE0R0ExVWREd0VCL3dRRUF3SUZvREFnQmdOVkhTVUJBZjhFRmpBVUJnZ3JCZ0VGQlFjREFRWUkKS3dZQkJRVUhBd0l3UGdZRFZSMFJCRGN3TllJSmJHOWpZV3hvYjNOMGdpSm5ieTFoZFhSb01DNWhjbU5vYVhSbApZM1IxY21VdVkyRnlZV1JvY21GekxtbHZod1IvQUFBQk1CMEdBMVVkRGdRV0JCVDZDSHpYQ2JzSUUyNkVaeEd0CnBpc2syZkNhMmpDQm9RWURWUjBqQklHWk1JR1dvWDZrZkRCNk1Tc3dLUVlEVlFRRERDSm5ieTFoZFhSb01DNWgKY21Ob2FYUmxZM1IxY21VdVkyRnlZV1JvY21GekxtbHZNUTB3Q3dZRFZRUUtEQVJrYjJOck1SVXdFd1lEVlFRTApEQXhoY21Ob2FYUmxZM1IxY21VeEN6QUpCZ05WQkFZVEFrSlNNUXN3Q1FZRFZRUUlEQUpUVURFTE1Ba0dBMVVFCkJ3d0NVMUNDRkQrZ2pzSVlRL2tEUFl1ek5tUDRtM2JQeE1zTk1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQ0FRQW4KOWtSSzd0M0wrQjNRdmp4SlhIVzg0SGhubjFPbGhRTVRHV24yNkJ1NjMxcGdJSzlqUWFHU2tXU3RUV0UyNFhWQgpBZkRsc3dMTUdUTW1QMDBRdDNDaTk2aTBocXN1djRyc1hhZDlGblZBZEJBa2lpbUc0TDMyM2NTb1U0ajNXMTNPCmNEcnFrZVY1S0NyQ3BtcWZ0bEg2dk12andBYS94K29FZDlJWGtsM0s5NzhqQ2QwcTNKZTVQSllDLytSVnVxN2QKNnM2Q2RlOGQ1Mk1WYnlqTkVYUEdIQ1pBcm43YmduUHBPNDR1U2NLM1dpbzNBTVorN3ZxcTFhTkhWOUlHZDRHTQprQkcydnRXbHB0RThUcDhoM21zQmlOSm5OMUFqZFFjOVA3bmxONFJxYUdmeU5pYmluMWQ4L2MzS0VOcU9acHRVClZ5WEtuak1BbGNYL1JKTmc3RGVZK2xDZ0Y4UHk4OVQyTXN5VTBmRHVaSlRKVUNobG9iaHpOUUhMSlNJdUpNbDYKMHdaelR6OXBxOTVOdERiaVRXSDFSNE1xcWdLRFlNQ0o0MUU4NjFIQ1NReUR3bmYvQTJ6NjB0S3BVK0pZWit4SQpNc21wekIwYW5YYXY1SEV6L1ZtcUZaclN2aGdUenpuZHVjQTYzaWpwalRsbkxkd2VaOXJMalhycU9GUzJ2UGtUCkRFZWlhcGc4OHB0SjNVbmc0QXY3ck95azlTaVhuNHc0ejREZlY1Y1hkL2VVcDJXWkgxV2R6Rm5LY1hwbFVJVkcKWE9rK2hyeWNOU2xaM0xvT1BGYWpDeVhDdXJSRHdkWmNyVDNZK1U0MmpLemdua2RtK3d0bXFRUkFIaC9jQi9ndgo3M29jOTY3Tmp5Ukc5UU9LZWdsNTBLSzc1U3BhbDVKVXNYbVFnaUdOdHc9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
        }