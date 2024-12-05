策略(增删改查)  version日期不要随意更改就固定是这个(旧版已过时 "2008-10-17")

北京MINIO_REGION: "cn-north-1" # aws 虽然可以自定义 , 建议使用aws一样的名称

增删改查 , 所有同权限arn:aws:s3:::后面是权限路径
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:PutObject",
                "s3:DeleteObject",
                "s3:GetBucketLocation"
            ],
            "Resource": [
                "arn:aws:s3:::*"
            ]
        }
    ]
}
```