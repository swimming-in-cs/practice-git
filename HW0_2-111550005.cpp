#include<iostream>
#include<string>
using namespace std;

int main(){
	int n;
	cin >> n;
	int num[100001] = {0};
	int a;
	for (int i = 0;i < n;i++){
		cin >> a;
		num[a]++;
	}
	int most = num[0];
	int f = 0;
	for (int i = 0;i < n;i++){
		if (num[i] >= f){
			most = i;
			f = num[i];
		}
	}
	cout << most << " " << f;
	return 0;
}
