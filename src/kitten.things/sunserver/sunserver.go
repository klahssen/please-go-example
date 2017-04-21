package main

import "fmt"
import "sunprotos/protos"

func main() {
	fmt.Println("sunserver!")
	fmt.Println(protos.UseItemXpBoostResponse_Result_name[1])
	a := protos.ListAvatarCustomizationsResponse{}
	customization := protos.AvatarCustomization{
		Labels: []protos.AvatarCustomization_Label{protos.AvatarCustomization_VIEWED, protos.AvatarCustomization_UNLOCKABLE},
	}
	a.AvatarCustomizations = &customization
	a.Result = protos.ListAvatarCustomizationsResponse_SUCCESS
	fmt.Println(a)
}
