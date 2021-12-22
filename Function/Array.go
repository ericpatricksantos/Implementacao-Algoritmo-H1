package Function

import "main/Model"

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func AuxH1(cluster1 []string, cluster2 []string) (bool, []string) {
	for _, item := range cluster1 {
		confirm := Contains(cluster2, item)

		if confirm {
			return true, UnionArray(cluster1, cluster2)
		}
	}

	return false, []string{}
}

func RemoveDuplicados(lista []string) ([]string, int) {
	var temp []string

	for _, x := range lista {
		if !Contains(temp, x) && len(x) > 0 {
			temp = append(temp, x)
		}
	}

	return temp, len(temp)
}

func UnionArray(input []string, out []string) []string {
	return append(input, out...)
}

func RemoveCluster(hash string, clusters []Model.Cluster) (result []Model.Cluster) {

	for _, item := range clusters {
		if item.Hash != hash {
			result = append(result, item)
		}
	}

	return result
}
