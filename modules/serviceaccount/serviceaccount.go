package serviceaccount

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/randomtoy/pulumi-yandex-cloud/modules/config"

	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
)

func bindingFolderMember(ctx *pulumi.Context, ya *config.Yandex, sa *config.ServiceAccount) error {
	pulumi.Sprintf("sa: %+v \r\n", sa)
	line := pulumi.Sprintf("serviceAccount:%s", sa.Id)
	member := pulumi.StringArray{line}.ToStringArrayOutput()
	pulumi.Sprintf("data: %+v", member)
	for i := range sa.Roles {
		_, err := yandex.NewResourcemanagerFolderIamBinding(ctx, sa.Name, &yandex.ResourcemanagerFolderIamBindingArgs{
			FolderId: pulumi.String(ya.FolderId),
			Members:  member,
			Role:     pulumi.String(sa.Roles[i]),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSA(ctx *pulumi.Context, ya *config.Yandex, sa *config.ServiceAccount) error {
	s, err := yandex.NewIamServiceAccount(ctx, sa.Name, &yandex.IamServiceAccountArgs{
		FolderId:    pulumi.String(ya.FolderId),
		Description: pulumi.String(sa.Description),
		Name:        pulumi.String(sa.Name),
	})
	if err != nil {
		return err
	}
	sid := s.ID().ToStringOutput()

	sa.Id = sid

	err = bindingFolderMember(ctx, ya, sa)
	ctx.Export("ServiceAccount ID: ", sid)

	return nil
}
