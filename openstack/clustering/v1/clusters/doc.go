/*
Package clusters provides information and interaction with the clusters through
the OpenStack Clustering service.

Example to Create a cluster

	createOpts := clusters.CreateOpts{
		Name:            "test-cluster",
		DesiredCapacity: 1,
		ProfileID:       "b7b870ee-d3c5-4a93-b9d7-846c53b2c2da",
	}

	cluster, err := clusters.Create(serviceClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Get Clusters

	clusterName := "cluster123"
	cluster, err := clusters.Get(serviceClient, clusterName).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cluster)

Example to List Clusters

	listOpts := clusters.ListOpts{
		Name: "testcluster",
	}

	allPages, err := clusters.ListDetail(serviceClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allClusters, err := clusters.ExtractClusters(allPages)
	if err != nil {
		panic(err)
	}

	for _, cluster := range allClusters {
		fmt.Printf("%+v\n", cluster)
	}

Example to Update a cluster

	updateOpts := clusters.UpdateOpts{
		Name:       "testcluster",
		ProfileID:  "b7b870ee-d3c5-4a93-b9d7-846c53b2c2da",
	}

	clusterID := "7d85f602-a948-4a30-afd4-e84f47471c15"
	cluster, err := clusters.Update(serviceClient, clusterName, opts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cluster)

Example to Delete a cluster

	clusterID := "dc6d336e3fc4c0a951b5698cd1236ee"
	err := clusters.Delete(serviceClient, clusterID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Resize a cluster

	number := 1
	maxSize := 5
	minSize := 1
	minStep := 1
	strict := true

	resizeOpts := clusters.ResizeOpts{
		AdjustmentType: clusters.ChangeInCapacityAdjustment,
		Number:         number,
		MaxSize:        &maxSize,
		MinSize:        &minSize,
		MinStep:        &minStep,
		Strict:         &strict,
	}

	actionID, err := clusters.Resize(client, clusterName, resizeOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to resize cluster: %v", err)
	}
	fmt.Println("Resize actionID", actionID)

Example to ScaleIn a cluster

	count := 2
	scaleInOpts := clusters.ScaleInOpts{
		Count: &count,
	}
	clusterID:  "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"

	action, err := clusters.ScaleIn(computeClient, clusterID, scaleInOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to list Policies for a Cluster

	clusterID := "7d85f602-a948-4a30-afd4-e84f47471c15"
	allPages, err := clusters.ListPolicies(serviceClient, clusterID, nil).AllPages()
	if err != nil {
		panic(err)
	}

	allClusterPolicies, err := clusters.ExtractClusterPolicies(allPages)
	if err != nil {
		panic(err)
	}

	for _, clusterPolicy := range allClusterPolicies {
		fmt.Printf("%+v\n", clusterPolicy)
	}

Example to Get a Cluster Policy

	clusterID := "7d85f602-a948-4a30-afd4-e84f47471c15"
	profileID := "714fe676-a08f-4196-b7af-61d52eeded15"
	clusterPolicy, err := clusterpolicies.Get(serviceCLient, clusterID, profileID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", clusterPolicy)

Example to Recover a cluster

	check := true
	checkCapacity := true
	recoverOpts := clusters.RecoverOpts{
		Operation:     clusters.RebuildRecovery,
		Check:         &check,
		CheckCapacity: &checkCapacity,
	}

	clusterID := "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"
	actionID, err := clusters.Recover(computeClient, clusterID, recoverOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("action=", actionID)

Example to attach a policy to a cluster

	enabled := true
	attachPolicyOpts := clusters.AttachPolicyOpts{
		PolicyID: "policy-123",
		Enabled:  &enabled,
	}
	clusterID := "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"
	actionID, err := clusters.AttachPolicy(serviceClient, clusterID, attachPolicyOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("Attach Policy actionID", actionID)

Example to update a policy to a cluster

	enabled := true
	updatePolicyOpts := clusters.UpdatePolicyOpts{
		PolicyID: "policy-123",
		Enabled:  &enabled,
	}
	clusterID := "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"
	actionID, err := clusters.AttachPolicy(serviceClient, clusterID, attachPolicyOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("Attach Policy actionID", actionID)

Example to detach a policy to cluster

	detachpolicyOpts := clusters.DetachPolicyOpts{
		PolicyID: "policy-123",
	}
	clusterID :=  "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"
	actionID, err := clusters.DetachPolicy(serviceClient, clusterID, detachpolicyOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("DetachPolicy actionID", actionID)
	actionID, err := clusters.UpdatePolicy(serviceClient, clusterID, updatePolicyOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Policy actionID", actionID)

Example to ScaleOut a cluster

	scaleOutOpts := clusters.ScaleOutOpts{
		Count: 2,
	}
	clusterID := "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"

	actionID, err := clusters.ScaleOut(computeClient, clusterID, scaleOutOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Check a cluster

	clusterID :=  "b7b870e3-d3c5-4a93-b9d7-846c53b2c2da"
	action, err := clusters.Check(computeClient, clusterID, clusters.CheckOpts{}).Extract()
	if err != nil {
		panic(err)
	}

*/
package clusters
