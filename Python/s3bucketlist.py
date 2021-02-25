#Import the boto3 module for AWS S3
import boto3

#Lists all of the s3 buckets in your account
#create variable
s3 = boto3.resource('s3')
#for loop
for bucket in s3.buckets.all():
    #print the list
    print (bucket.name)
