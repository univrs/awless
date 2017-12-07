/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
// This file was automatically generated with go generate
package awsspec

var APIPerTemplateDefName = map[string]string{
	"attachalarm":               "cloudwatch",
	"attachcontainertask":       "ecs",
	"attachelasticip":           "ec2",
	"attachinstance":            "elbv2",
	"attachinstanceprofile":     "ec2",
	"attachinternetgateway":     "ec2",
	"attachmfadevice":           "iam",
	"attachnetworkinterface":    "ec2",
	"attachpolicy":              "iam",
	"attachrole":                "iam",
	"attachroutetable":          "ec2",
	"attachsecuritygroup":       "ec2",
	"attachuser":                "iam",
	"attachvolume":              "ec2",
	"authenticateregistry":      "ecr",
	"checkcertificate":          "acm",
	"checkdatabase":             "rds",
	"checkdistribution":         "cloudfront",
	"checkinstance":             "ec2",
	"checkloadbalancer":         "elbv2",
	"checknatgateway":           "ec2",
	"checknetworkinterface":     "ec2",
	"checkscalinggroup":         "autoscaling",
	"checksecuritygroup":        "ec2",
	"checkvolume":               "ec2",
	"copyimage":                 "ec2",
	"copysnapshot":              "ec2",
	"createaccesskey":           "iam",
	"createalarm":               "cloudwatch",
	"createappscalingpolicy":    "applicationautoscaling",
	"createappscalingtarget":    "applicationautoscaling",
	"createbucket":              "s3",
	"createcertificate":         "acm",
	"createcontainercluster":    "ecs",
	"createdatabase":            "rds",
	"createdbsubnetgroup":       "rds",
	"createdistribution":        "cloudfront",
	"createelasticip":           "ec2",
	"createfunction":            "lambda",
	"creategroup":               "iam",
	"createimage":               "ec2",
	"createinstance":            "ec2",
	"createinstanceprofile":     "iam",
	"createinternetgateway":     "ec2",
	"createkeypair":             "ec2",
	"createlaunchconfiguration": "autoscaling",
	"createlistener":            "elbv2",
	"createloadbalancer":        "elbv2",
	"createloginprofile":        "iam",
	"createmfadevice":           "iam",
	"createnatgateway":          "ec2",
	"createnetworkinterface":    "ec2",
	"createpolicy":              "iam",
	"createqueue":               "sqs",
	"createrecord":              "route53",
	"createrepository":          "ecr",
	"createrole":                "iam",
	"createroute":               "ec2",
	"createroutetable":          "ec2",
	"creates3object":            "s3",
	"createscalinggroup":        "autoscaling",
	"createscalingpolicy":       "autoscaling",
	"createsecuritygroup":       "ec2",
	"createsnapshot":            "ec2",
	"createstack":               "cloudformation",
	"createsubnet":              "ec2",
	"createsubscription":        "sns",
	"createtag":                 "ec2",
	"createtargetgroup":         "elbv2",
	"createtopic":               "sns",
	"createuser":                "iam",
	"createvolume":              "ec2",
	"createvpc":                 "ec2",
	"createzone":                "route53",
	"deleteaccesskey":           "iam",
	"deletealarm":               "cloudwatch",
	"deleteappscalingpolicy":    "applicationautoscaling",
	"deleteappscalingtarget":    "applicationautoscaling",
	"deletebucket":              "s3",
	"deletecertificate":         "acm",
	"deletecontainercluster":    "ecs",
	"deletecontainertask":       "ecs",
	"deletedatabase":            "rds",
	"deletedbsubnetgroup":       "rds",
	"deletedistribution":        "cloudfront",
	"deleteelasticip":           "ec2",
	"deletefunction":            "lambda",
	"deletegroup":               "iam",
	"deleteimage":               "ec2",
	"deleteinstance":            "ec2",
	"deleteinstanceprofile":     "iam",
	"deleteinternetgateway":     "ec2",
	"deletekeypair":             "ec2",
	"deletelaunchconfiguration": "autoscaling",
	"deletelistener":            "elbv2",
	"deleteloadbalancer":        "elbv2",
	"deleteloginprofile":        "iam",
	"deletemfadevice":           "iam",
	"deletenatgateway":          "ec2",
	"deletenetworkinterface":    "ec2",
	"deletepolicy":              "iam",
	"deletequeue":               "sqs",
	"deleterecord":              "route53",
	"deleterepository":          "ecr",
	"deleterole":                "iam",
	"deleteroute":               "ec2",
	"deleteroutetable":          "ec2",
	"deletes3object":            "s3",
	"deletescalinggroup":        "autoscaling",
	"deletescalingpolicy":       "autoscaling",
	"deletesecuritygroup":       "ec2",
	"deletesnapshot":            "ec2",
	"deletestack":               "cloudformation",
	"deletesubnet":              "ec2",
	"deletesubscription":        "sns",
	"deletetag":                 "ec2",
	"deletetargetgroup":         "elbv2",
	"deletetopic":               "sns",
	"deleteuser":                "iam",
	"deletevolume":              "ec2",
	"deletevpc":                 "ec2",
	"deletezone":                "route53",
	"detachalarm":               "cloudwatch",
	"detachcontainertask":       "ecs",
	"detachelasticip":           "ec2",
	"detachinstance":            "elbv2",
	"detachinstanceprofile":     "ec2",
	"detachinternetgateway":     "ec2",
	"detachmfadevice":           "iam",
	"detachnetworkinterface":    "ec2",
	"detachpolicy":              "iam",
	"detachrole":                "iam",
	"detachroutetable":          "ec2",
	"detachsecuritygroup":       "ec2",
	"detachuser":                "iam",
	"detachvolume":              "ec2",
	"importimage":               "ec2",
	"startalarm":                "cloudwatch",
	"startcontainertask":        "ecs",
	"startdatabase":             "rds",
	"startinstance":             "ec2",
	"stopalarm":                 "cloudwatch",
	"stopcontainertask":         "ecs",
	"stopdatabase":              "rds",
	"stopinstance":              "ec2",
	"updatebucket":              "s3",
	"updatecontainertask":       "ecs",
	"updatedistribution":        "cloudfront",
	"updateimage":               "ec2",
	"updateinstance":            "ec2",
	"updateloginprofile":        "iam",
	"updatepolicy":              "iam",
	"updaterecord":              "route53",
	"updates3object":            "s3",
	"updatescalinggroup":        "autoscaling",
	"updatesecuritygroup":       "ec2",
	"updatestack":               "cloudformation",
	"updatesubnet":              "ec2",
	"updatetargetgroup":         "elbv2",
}

var AWSTemplatesDefinitions = map[string]Definition{
	"attachalarm": {
		Action: "attach",
		Entity: "alarm",
		Api:    "cloudwatch",
		Params: new(AttachAlarm).Params(),
	},
	"attachcontainertask": {
		Action: "attach",
		Entity: "containertask",
		Api:    "ecs",
		Params: new(AttachContainertask).Params(),
	},
	"attachelasticip": {
		Action: "attach",
		Entity: "elasticip",
		Api:    "ec2",
		Params: new(AttachElasticip).Params(),
	},
	"attachinstance": {
		Action: "attach",
		Entity: "instance",
		Api:    "elbv2",
		Params: new(AttachInstance).Params(),
	},
	"attachinstanceprofile": {
		Action: "attach",
		Entity: "instanceprofile",
		Api:    "ec2",
		Params: new(AttachInstanceprofile).Params(),
	},
	"attachinternetgateway": {
		Action: "attach",
		Entity: "internetgateway",
		Api:    "ec2",
		Params: new(AttachInternetgateway).Params(),
	},
	"attachmfadevice": {
		Action: "attach",
		Entity: "mfadevice",
		Api:    "iam",
		Params: new(AttachMfadevice).Params(),
	},
	"attachnetworkinterface": {
		Action: "attach",
		Entity: "networkinterface",
		Api:    "ec2",
		Params: new(AttachNetworkinterface).Params(),
	},
	"attachpolicy": {
		Action: "attach",
		Entity: "policy",
		Api:    "iam",
		Params: new(AttachPolicy).Params(),
	},
	"attachrole": {
		Action: "attach",
		Entity: "role",
		Api:    "iam",
		Params: new(AttachRole).Params(),
	},
	"attachroutetable": {
		Action: "attach",
		Entity: "routetable",
		Api:    "ec2",
		Params: new(AttachRoutetable).Params(),
	},
	"attachsecuritygroup": {
		Action: "attach",
		Entity: "securitygroup",
		Api:    "ec2",
		Params: new(AttachSecuritygroup).Params(),
	},
	"attachuser": {
		Action: "attach",
		Entity: "user",
		Api:    "iam",
		Params: new(AttachUser).Params(),
	},
	"attachvolume": {
		Action: "attach",
		Entity: "volume",
		Api:    "ec2",
		Params: new(AttachVolume).Params(),
	},
	"authenticateregistry": {
		Action: "authenticate",
		Entity: "registry",
		Api:    "ecr",
		Params: new(AuthenticateRegistry).Params(),
	},
	"checkcertificate": {
		Action: "check",
		Entity: "certificate",
		Api:    "acm",
		Params: new(CheckCertificate).Params(),
	},
	"checkdatabase": {
		Action: "check",
		Entity: "database",
		Api:    "rds",
		Params: new(CheckDatabase).Params(),
	},
	"checkdistribution": {
		Action: "check",
		Entity: "distribution",
		Api:    "cloudfront",
		Params: new(CheckDistribution).Params(),
	},
	"checkinstance": {
		Action: "check",
		Entity: "instance",
		Api:    "ec2",
		Params: new(CheckInstance).Params(),
	},
	"checkloadbalancer": {
		Action: "check",
		Entity: "loadbalancer",
		Api:    "elbv2",
		Params: new(CheckLoadbalancer).Params(),
	},
	"checknatgateway": {
		Action: "check",
		Entity: "natgateway",
		Api:    "ec2",
		Params: new(CheckNatgateway).Params(),
	},
	"checknetworkinterface": {
		Action: "check",
		Entity: "networkinterface",
		Api:    "ec2",
		Params: new(CheckNetworkinterface).Params(),
	},
	"checkscalinggroup": {
		Action: "check",
		Entity: "scalinggroup",
		Api:    "autoscaling",
		Params: new(CheckScalinggroup).Params(),
	},
	"checksecuritygroup": {
		Action: "check",
		Entity: "securitygroup",
		Api:    "ec2",
		Params: new(CheckSecuritygroup).Params(),
	},
	"checkvolume": {
		Action: "check",
		Entity: "volume",
		Api:    "ec2",
		Params: new(CheckVolume).Params(),
	},
	"copyimage": {
		Action: "copy",
		Entity: "image",
		Api:    "ec2",
		Params: new(CopyImage).Params(),
	},
	"copysnapshot": {
		Action: "copy",
		Entity: "snapshot",
		Api:    "ec2",
		Params: new(CopySnapshot).Params(),
	},
	"createaccesskey": {
		Action: "create",
		Entity: "accesskey",
		Api:    "iam",
		Params: new(CreateAccesskey).Params(),
	},
	"createalarm": {
		Action: "create",
		Entity: "alarm",
		Api:    "cloudwatch",
		Params: new(CreateAlarm).Params(),
	},
	"createappscalingpolicy": {
		Action: "create",
		Entity: "appscalingpolicy",
		Api:    "applicationautoscaling",
		Params: new(CreateAppscalingpolicy).Params(),
	},
	"createappscalingtarget": {
		Action: "create",
		Entity: "appscalingtarget",
		Api:    "applicationautoscaling",
		Params: new(CreateAppscalingtarget).Params(),
	},
	"createbucket": {
		Action: "create",
		Entity: "bucket",
		Api:    "s3",
		Params: new(CreateBucket).Params(),
	},
	"createcertificate": {
		Action: "create",
		Entity: "certificate",
		Api:    "acm",
		Params: new(CreateCertificate).Params(),
	},
	"createcontainercluster": {
		Action: "create",
		Entity: "containercluster",
		Api:    "ecs",
		Params: new(CreateContainercluster).Params(),
	},
	"createdatabase": {
		Action: "create",
		Entity: "database",
		Api:    "rds",
		Params: new(CreateDatabase).Params(),
	},
	"createdbsubnetgroup": {
		Action: "create",
		Entity: "dbsubnetgroup",
		Api:    "rds",
		Params: new(CreateDbsubnetgroup).Params(),
	},
	"createdistribution": {
		Action: "create",
		Entity: "distribution",
		Api:    "cloudfront",
		Params: new(CreateDistribution).Params(),
	},
	"createelasticip": {
		Action: "create",
		Entity: "elasticip",
		Api:    "ec2",
		Params: new(CreateElasticip).Params(),
	},
	"createfunction": {
		Action: "create",
		Entity: "function",
		Api:    "lambda",
		Params: new(CreateFunction).Params(),
	},
	"creategroup": {
		Action: "create",
		Entity: "group",
		Api:    "iam",
		Params: new(CreateGroup).Params(),
	},
	"createimage": {
		Action: "create",
		Entity: "image",
		Api:    "ec2",
		Params: new(CreateImage).Params(),
	},
	"createinstance": {
		Action: "create",
		Entity: "instance",
		Api:    "ec2",
		Params: new(CreateInstance).Params(),
	},
	"createinstanceprofile": {
		Action: "create",
		Entity: "instanceprofile",
		Api:    "iam",
		Params: new(CreateInstanceprofile).Params(),
	},
	"createinternetgateway": {
		Action: "create",
		Entity: "internetgateway",
		Api:    "ec2",
		Params: new(CreateInternetgateway).Params(),
	},
	"createkeypair": {
		Action: "create",
		Entity: "keypair",
		Api:    "ec2",
		Params: new(CreateKeypair).Params(),
	},
	"createlaunchconfiguration": {
		Action: "create",
		Entity: "launchconfiguration",
		Api:    "autoscaling",
		Params: new(CreateLaunchconfiguration).Params(),
	},
	"createlistener": {
		Action: "create",
		Entity: "listener",
		Api:    "elbv2",
		Params: new(CreateListener).Params(),
	},
	"createloadbalancer": {
		Action: "create",
		Entity: "loadbalancer",
		Api:    "elbv2",
		Params: new(CreateLoadbalancer).Params(),
	},
	"createloginprofile": {
		Action: "create",
		Entity: "loginprofile",
		Api:    "iam",
		Params: new(CreateLoginprofile).Params(),
	},
	"createmfadevice": {
		Action: "create",
		Entity: "mfadevice",
		Api:    "iam",
		Params: new(CreateMfadevice).Params(),
	},
	"createnatgateway": {
		Action: "create",
		Entity: "natgateway",
		Api:    "ec2",
		Params: new(CreateNatgateway).Params(),
	},
	"createnetworkinterface": {
		Action: "create",
		Entity: "networkinterface",
		Api:    "ec2",
		Params: new(CreateNetworkinterface).Params(),
	},
	"createpolicy": {
		Action: "create",
		Entity: "policy",
		Api:    "iam",
		Params: new(CreatePolicy).Params(),
	},
	"createqueue": {
		Action: "create",
		Entity: "queue",
		Api:    "sqs",
		Params: new(CreateQueue).Params(),
	},
	"createrecord": {
		Action: "create",
		Entity: "record",
		Api:    "route53",
		Params: new(CreateRecord).Params(),
	},
	"createrepository": {
		Action: "create",
		Entity: "repository",
		Api:    "ecr",
		Params: new(CreateRepository).Params(),
	},
	"createrole": {
		Action: "create",
		Entity: "role",
		Api:    "iam",
		Params: new(CreateRole).Params(),
	},
	"createroute": {
		Action: "create",
		Entity: "route",
		Api:    "ec2",
		Params: new(CreateRoute).Params(),
	},
	"createroutetable": {
		Action: "create",
		Entity: "routetable",
		Api:    "ec2",
		Params: new(CreateRoutetable).Params(),
	},
	"creates3object": {
		Action: "create",
		Entity: "s3object",
		Api:    "s3",
		Params: new(CreateS3object).Params(),
	},
	"createscalinggroup": {
		Action: "create",
		Entity: "scalinggroup",
		Api:    "autoscaling",
		Params: new(CreateScalinggroup).Params(),
	},
	"createscalingpolicy": {
		Action: "create",
		Entity: "scalingpolicy",
		Api:    "autoscaling",
		Params: new(CreateScalingpolicy).Params(),
	},
	"createsecuritygroup": {
		Action: "create",
		Entity: "securitygroup",
		Api:    "ec2",
		Params: new(CreateSecuritygroup).Params(),
	},
	"createsnapshot": {
		Action: "create",
		Entity: "snapshot",
		Api:    "ec2",
		Params: new(CreateSnapshot).Params(),
	},
	"createstack": {
		Action: "create",
		Entity: "stack",
		Api:    "cloudformation",
		Params: new(CreateStack).Params(),
	},
	"createsubnet": {
		Action: "create",
		Entity: "subnet",
		Api:    "ec2",
		Params: new(CreateSubnet).Params(),
	},
	"createsubscription": {
		Action: "create",
		Entity: "subscription",
		Api:    "sns",
		Params: new(CreateSubscription).Params(),
	},
	"createtag": {
		Action: "create",
		Entity: "tag",
		Api:    "ec2",
		Params: new(CreateTag).Params(),
	},
	"createtargetgroup": {
		Action: "create",
		Entity: "targetgroup",
		Api:    "elbv2",
		Params: new(CreateTargetgroup).Params(),
	},
	"createtopic": {
		Action: "create",
		Entity: "topic",
		Api:    "sns",
		Params: new(CreateTopic).Params(),
	},
	"createuser": {
		Action: "create",
		Entity: "user",
		Api:    "iam",
		Params: new(CreateUser).Params(),
	},
	"createvolume": {
		Action: "create",
		Entity: "volume",
		Api:    "ec2",
		Params: new(CreateVolume).Params(),
	},
	"createvpc": {
		Action: "create",
		Entity: "vpc",
		Api:    "ec2",
		Params: new(CreateVpc).Params(),
	},
	"createzone": {
		Action: "create",
		Entity: "zone",
		Api:    "route53",
		Params: new(CreateZone).Params(),
	},
	"deleteaccesskey": {
		Action: "delete",
		Entity: "accesskey",
		Api:    "iam",
		Params: new(DeleteAccesskey).Params(),
	},
	"deletealarm": {
		Action: "delete",
		Entity: "alarm",
		Api:    "cloudwatch",
		Params: new(DeleteAlarm).Params(),
	},
	"deleteappscalingpolicy": {
		Action: "delete",
		Entity: "appscalingpolicy",
		Api:    "applicationautoscaling",
		Params: new(DeleteAppscalingpolicy).Params(),
	},
	"deleteappscalingtarget": {
		Action: "delete",
		Entity: "appscalingtarget",
		Api:    "applicationautoscaling",
		Params: new(DeleteAppscalingtarget).Params(),
	},
	"deletebucket": {
		Action: "delete",
		Entity: "bucket",
		Api:    "s3",
		Params: new(DeleteBucket).Params(),
	},
	"deletecertificate": {
		Action: "delete",
		Entity: "certificate",
		Api:    "acm",
		Params: new(DeleteCertificate).Params(),
	},
	"deletecontainercluster": {
		Action: "delete",
		Entity: "containercluster",
		Api:    "ecs",
		Params: new(DeleteContainercluster).Params(),
	},
	"deletecontainertask": {
		Action: "delete",
		Entity: "containertask",
		Api:    "ecs",
		Params: new(DeleteContainertask).Params(),
	},
	"deletedatabase": {
		Action: "delete",
		Entity: "database",
		Api:    "rds",
		Params: new(DeleteDatabase).Params(),
	},
	"deletedbsubnetgroup": {
		Action: "delete",
		Entity: "dbsubnetgroup",
		Api:    "rds",
		Params: new(DeleteDbsubnetgroup).Params(),
	},
	"deletedistribution": {
		Action: "delete",
		Entity: "distribution",
		Api:    "cloudfront",
		Params: new(DeleteDistribution).Params(),
	},
	"deleteelasticip": {
		Action: "delete",
		Entity: "elasticip",
		Api:    "ec2",
		Params: new(DeleteElasticip).Params(),
	},
	"deletefunction": {
		Action: "delete",
		Entity: "function",
		Api:    "lambda",
		Params: new(DeleteFunction).Params(),
	},
	"deletegroup": {
		Action: "delete",
		Entity: "group",
		Api:    "iam",
		Params: new(DeleteGroup).Params(),
	},
	"deleteimage": {
		Action: "delete",
		Entity: "image",
		Api:    "ec2",
		Params: new(DeleteImage).Params(),
	},
	"deleteinstance": {
		Action: "delete",
		Entity: "instance",
		Api:    "ec2",
		Params: new(DeleteInstance).Params(),
	},
	"deleteinstanceprofile": {
		Action: "delete",
		Entity: "instanceprofile",
		Api:    "iam",
		Params: new(DeleteInstanceprofile).Params(),
	},
	"deleteinternetgateway": {
		Action: "delete",
		Entity: "internetgateway",
		Api:    "ec2",
		Params: new(DeleteInternetgateway).Params(),
	},
	"deletekeypair": {
		Action: "delete",
		Entity: "keypair",
		Api:    "ec2",
		Params: new(DeleteKeypair).Params(),
	},
	"deletelaunchconfiguration": {
		Action: "delete",
		Entity: "launchconfiguration",
		Api:    "autoscaling",
		Params: new(DeleteLaunchconfiguration).Params(),
	},
	"deletelistener": {
		Action: "delete",
		Entity: "listener",
		Api:    "elbv2",
		Params: new(DeleteListener).Params(),
	},
	"deleteloadbalancer": {
		Action: "delete",
		Entity: "loadbalancer",
		Api:    "elbv2",
		Params: new(DeleteLoadbalancer).Params(),
	},
	"deleteloginprofile": {
		Action: "delete",
		Entity: "loginprofile",
		Api:    "iam",
		Params: new(DeleteLoginprofile).Params(),
	},
	"deletemfadevice": {
		Action: "delete",
		Entity: "mfadevice",
		Api:    "iam",
		Params: new(DeleteMfadevice).Params(),
	},
	"deletenatgateway": {
		Action: "delete",
		Entity: "natgateway",
		Api:    "ec2",
		Params: new(DeleteNatgateway).Params(),
	},
	"deletenetworkinterface": {
		Action: "delete",
		Entity: "networkinterface",
		Api:    "ec2",
		Params: new(DeleteNetworkinterface).Params(),
	},
	"deletepolicy": {
		Action: "delete",
		Entity: "policy",
		Api:    "iam",
		Params: new(DeletePolicy).Params(),
	},
	"deletequeue": {
		Action: "delete",
		Entity: "queue",
		Api:    "sqs",
		Params: new(DeleteQueue).Params(),
	},
	"deleterecord": {
		Action: "delete",
		Entity: "record",
		Api:    "route53",
		Params: new(DeleteRecord).Params(),
	},
	"deleterepository": {
		Action: "delete",
		Entity: "repository",
		Api:    "ecr",
		Params: new(DeleteRepository).Params(),
	},
	"deleterole": {
		Action: "delete",
		Entity: "role",
		Api:    "iam",
		Params: new(DeleteRole).Params(),
	},
	"deleteroute": {
		Action: "delete",
		Entity: "route",
		Api:    "ec2",
		Params: new(DeleteRoute).Params(),
	},
	"deleteroutetable": {
		Action: "delete",
		Entity: "routetable",
		Api:    "ec2",
		Params: new(DeleteRoutetable).Params(),
	},
	"deletes3object": {
		Action: "delete",
		Entity: "s3object",
		Api:    "s3",
		Params: new(DeleteS3object).Params(),
	},
	"deletescalinggroup": {
		Action: "delete",
		Entity: "scalinggroup",
		Api:    "autoscaling",
		Params: new(DeleteScalinggroup).Params(),
	},
	"deletescalingpolicy": {
		Action: "delete",
		Entity: "scalingpolicy",
		Api:    "autoscaling",
		Params: new(DeleteScalingpolicy).Params(),
	},
	"deletesecuritygroup": {
		Action: "delete",
		Entity: "securitygroup",
		Api:    "ec2",
		Params: new(DeleteSecuritygroup).Params(),
	},
	"deletesnapshot": {
		Action: "delete",
		Entity: "snapshot",
		Api:    "ec2",
		Params: new(DeleteSnapshot).Params(),
	},
	"deletestack": {
		Action: "delete",
		Entity: "stack",
		Api:    "cloudformation",
		Params: new(DeleteStack).Params(),
	},
	"deletesubnet": {
		Action: "delete",
		Entity: "subnet",
		Api:    "ec2",
		Params: new(DeleteSubnet).Params(),
	},
	"deletesubscription": {
		Action: "delete",
		Entity: "subscription",
		Api:    "sns",
		Params: new(DeleteSubscription).Params(),
	},
	"deletetag": {
		Action: "delete",
		Entity: "tag",
		Api:    "ec2",
		Params: new(DeleteTag).Params(),
	},
	"deletetargetgroup": {
		Action: "delete",
		Entity: "targetgroup",
		Api:    "elbv2",
		Params: new(DeleteTargetgroup).Params(),
	},
	"deletetopic": {
		Action: "delete",
		Entity: "topic",
		Api:    "sns",
		Params: new(DeleteTopic).Params(),
	},
	"deleteuser": {
		Action: "delete",
		Entity: "user",
		Api:    "iam",
		Params: new(DeleteUser).Params(),
	},
	"deletevolume": {
		Action: "delete",
		Entity: "volume",
		Api:    "ec2",
		Params: new(DeleteVolume).Params(),
	},
	"deletevpc": {
		Action: "delete",
		Entity: "vpc",
		Api:    "ec2",
		Params: new(DeleteVpc).Params(),
	},
	"deletezone": {
		Action: "delete",
		Entity: "zone",
		Api:    "route53",
		Params: new(DeleteZone).Params(),
	},
	"detachalarm": {
		Action: "detach",
		Entity: "alarm",
		Api:    "cloudwatch",
		Params: new(DetachAlarm).Params(),
	},
	"detachcontainertask": {
		Action: "detach",
		Entity: "containertask",
		Api:    "ecs",
		Params: new(DetachContainertask).Params(),
	},
	"detachelasticip": {
		Action: "detach",
		Entity: "elasticip",
		Api:    "ec2",
		Params: new(DetachElasticip).Params(),
	},
	"detachinstance": {
		Action: "detach",
		Entity: "instance",
		Api:    "elbv2",
		Params: new(DetachInstance).Params(),
	},
	"detachinstanceprofile": {
		Action: "detach",
		Entity: "instanceprofile",
		Api:    "ec2",
		Params: new(DetachInstanceprofile).Params(),
	},
	"detachinternetgateway": {
		Action: "detach",
		Entity: "internetgateway",
		Api:    "ec2",
		Params: new(DetachInternetgateway).Params(),
	},
	"detachmfadevice": {
		Action: "detach",
		Entity: "mfadevice",
		Api:    "iam",
		Params: new(DetachMfadevice).Params(),
	},
	"detachnetworkinterface": {
		Action: "detach",
		Entity: "networkinterface",
		Api:    "ec2",
		Params: new(DetachNetworkinterface).Params(),
	},
	"detachpolicy": {
		Action: "detach",
		Entity: "policy",
		Api:    "iam",
		Params: new(DetachPolicy).Params(),
	},
	"detachrole": {
		Action: "detach",
		Entity: "role",
		Api:    "iam",
		Params: new(DetachRole).Params(),
	},
	"detachroutetable": {
		Action: "detach",
		Entity: "routetable",
		Api:    "ec2",
		Params: new(DetachRoutetable).Params(),
	},
	"detachsecuritygroup": {
		Action: "detach",
		Entity: "securitygroup",
		Api:    "ec2",
		Params: new(DetachSecuritygroup).Params(),
	},
	"detachuser": {
		Action: "detach",
		Entity: "user",
		Api:    "iam",
		Params: new(DetachUser).Params(),
	},
	"detachvolume": {
		Action: "detach",
		Entity: "volume",
		Api:    "ec2",
		Params: new(DetachVolume).Params(),
	},
	"importimage": {
		Action: "import",
		Entity: "image",
		Api:    "ec2",
		Params: new(ImportImage).Params(),
	},
	"startalarm": {
		Action: "start",
		Entity: "alarm",
		Api:    "cloudwatch",
		Params: new(StartAlarm).Params(),
	},
	"startcontainertask": {
		Action: "start",
		Entity: "containertask",
		Api:    "ecs",
		Params: new(StartContainertask).Params(),
	},
	"startdatabase": {
		Action: "start",
		Entity: "database",
		Api:    "rds",
		Params: new(StartDatabase).Params(),
	},
	"startinstance": {
		Action: "start",
		Entity: "instance",
		Api:    "ec2",
		Params: new(StartInstance).Params(),
	},
	"stopalarm": {
		Action: "stop",
		Entity: "alarm",
		Api:    "cloudwatch",
		Params: new(StopAlarm).Params(),
	},
	"stopcontainertask": {
		Action: "stop",
		Entity: "containertask",
		Api:    "ecs",
		Params: new(StopContainertask).Params(),
	},
	"stopdatabase": {
		Action: "stop",
		Entity: "database",
		Api:    "rds",
		Params: new(StopDatabase).Params(),
	},
	"stopinstance": {
		Action: "stop",
		Entity: "instance",
		Api:    "ec2",
		Params: new(StopInstance).Params(),
	},
	"updatebucket": {
		Action: "update",
		Entity: "bucket",
		Api:    "s3",
		Params: new(UpdateBucket).Params(),
	},
	"updatecontainertask": {
		Action: "update",
		Entity: "containertask",
		Api:    "ecs",
		Params: new(UpdateContainertask).Params(),
	},
	"updatedistribution": {
		Action: "update",
		Entity: "distribution",
		Api:    "cloudfront",
		Params: new(UpdateDistribution).Params(),
	},
	"updateimage": {
		Action: "update",
		Entity: "image",
		Api:    "ec2",
		Params: new(UpdateImage).Params(),
	},
	"updateinstance": {
		Action: "update",
		Entity: "instance",
		Api:    "ec2",
		Params: new(UpdateInstance).Params(),
	},
	"updateloginprofile": {
		Action: "update",
		Entity: "loginprofile",
		Api:    "iam",
		Params: new(UpdateLoginprofile).Params(),
	},
	"updatepolicy": {
		Action: "update",
		Entity: "policy",
		Api:    "iam",
		Params: new(UpdatePolicy).Params(),
	},
	"updaterecord": {
		Action: "update",
		Entity: "record",
		Api:    "route53",
		Params: new(UpdateRecord).Params(),
	},
	"updates3object": {
		Action: "update",
		Entity: "s3object",
		Api:    "s3",
		Params: new(UpdateS3object).Params(),
	},
	"updatescalinggroup": {
		Action: "update",
		Entity: "scalinggroup",
		Api:    "autoscaling",
		Params: new(UpdateScalinggroup).Params(),
	},
	"updatesecuritygroup": {
		Action: "update",
		Entity: "securitygroup",
		Api:    "ec2",
		Params: new(UpdateSecuritygroup).Params(),
	},
	"updatestack": {
		Action: "update",
		Entity: "stack",
		Api:    "cloudformation",
		Params: new(UpdateStack).Params(),
	},
	"updatesubnet": {
		Action: "update",
		Entity: "subnet",
		Api:    "ec2",
		Params: new(UpdateSubnet).Params(),
	},
	"updatetargetgroup": {
		Action: "update",
		Entity: "targetgroup",
		Api:    "elbv2",
		Params: new(UpdateTargetgroup).Params(),
	},
}

var DriverSupportedActions = map[string][]string{
	"attach":       {"alarm", "containertask", "elasticip", "instance", "instanceprofile", "internetgateway", "mfadevice", "networkinterface", "policy", "role", "routetable", "securitygroup", "user", "volume"},
	"authenticate": {"registry"},
	"check":        {"certificate", "database", "distribution", "instance", "loadbalancer", "natgateway", "networkinterface", "scalinggroup", "securitygroup", "volume"},
	"copy":         {"image", "snapshot"},
	"create":       {"accesskey", "alarm", "appscalingpolicy", "appscalingtarget", "bucket", "certificate", "containercluster", "database", "dbsubnetgroup", "distribution", "elasticip", "function", "group", "image", "instance", "instanceprofile", "internetgateway", "keypair", "launchconfiguration", "listener", "loadbalancer", "loginprofile", "mfadevice", "natgateway", "networkinterface", "policy", "queue", "record", "repository", "role", "route", "routetable", "s3object", "scalinggroup", "scalingpolicy", "securitygroup", "snapshot", "stack", "subnet", "subscription", "tag", "targetgroup", "topic", "user", "volume", "vpc", "zone"},
	"delete":       {"accesskey", "alarm", "appscalingpolicy", "appscalingtarget", "bucket", "certificate", "containercluster", "containertask", "database", "dbsubnetgroup", "distribution", "elasticip", "function", "group", "image", "instance", "instanceprofile", "internetgateway", "keypair", "launchconfiguration", "listener", "loadbalancer", "loginprofile", "mfadevice", "natgateway", "networkinterface", "policy", "queue", "record", "repository", "role", "route", "routetable", "s3object", "scalinggroup", "scalingpolicy", "securitygroup", "snapshot", "stack", "subnet", "subscription", "tag", "targetgroup", "topic", "user", "volume", "vpc", "zone"},
	"detach":       {"alarm", "containertask", "elasticip", "instance", "instanceprofile", "internetgateway", "mfadevice", "networkinterface", "policy", "role", "routetable", "securitygroup", "user", "volume"},
	"import":       {"image"},
	"start":        {"alarm", "containertask", "database", "instance"},
	"stop":         {"alarm", "containertask", "database", "instance"},
	"update":       {"bucket", "containertask", "distribution", "image", "instance", "loginprofile", "policy", "record", "s3object", "scalinggroup", "securitygroup", "stack", "subnet", "targetgroup"},
}
