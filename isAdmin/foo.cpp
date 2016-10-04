// copy and pasted from http://d.hatena.ne.jp/s-kita/201103

#include <windows.h>
#include <aclapi.h>
#include <stdio.h>

int main()
{
    PSECURITY_DESCRIPTOR pSD;
    DWORD nLengthNeeded = 0;
    char szAccountName[256];
    DWORD dwAccountNameSize;
    char szDomainName[256];
    DWORD dwDomainNameSize;
    SID_NAME_USE snu;
    BOOL bOwnerDefault;
    PSID pSID;

    GetKernelObjectSecurity(GetCurrentProcess(), OWNER_SECURITY_INFORMATION, NULL, nLengthNeeded, &nLengthNeeded);

    pSD = (PSECURITY_DESCRIPTOR)LocalAlloc(LPTR, nLengthNeeded);
    if (pSD == NULL) {
        return 1;
    }

    //1.GetKernelObjectSecurity関数で、カーネルオブジェクトのセキュリティ記述子を取得する
    GetKernelObjectSecurity(GetCurrentProcess(), OWNER_SECURITY_INFORMATION, pSD, nLengthNeeded, &nLengthNeeded);

    //2.GetSecurityDescriptorOwner関数で、セキュリティ記述子からSIDを取得する
    GetSecurityDescriptorOwner(pSD, &pSID, &bOwnerDefault);

    if (pSID == NULL) {
        puts("No owner");
        goto EXIT_FUNC;
    }


    dwAccountNameSize = sizeof(szAccountName)/sizeof(szAccountName[0]);
    dwDomainNameSize = sizeof(szDomainName)/sizeof(szDomainName[0]);

    //3.LookupAccountSid関数で、SIDに対するアカウント情報を取得する
    LookupAccountSidA(NULL,
        pSID,
        szAccountName,
        &dwAccountNameSize,
        szDomainName,
        &dwDomainNameSize,
        &snu);

    printf("AccountName: %s\n", szAccountName);
    printf("DomainName: %s\n", szDomainName);

    switch (snu) {
    case SidTypeUser:
        puts("SidTypeUser");
        break;
    case SidTypeGroup:
        puts("SidTypeGroup");
        break;
    case SidTypeDomain:
        puts("SidTypeDomain");
        break;
    case SidTypeAlias:
        puts("SidTypeAlias");
        break;
    case SidTypeWellKnownGroup:
        puts("SidTypeWellKnownGroup");
        break;
    case SidTypeDeletedAccount:
        puts("SidTypeDeletedAccount");
        break;
    case SidTypeInvalid:
        puts("SidTypeInvalid");
        break;
    case SidTypeUnknown:
        puts("SidTypeUnknown");
        break;
    case SidTypeComputer:
        puts("SidTypeComputer");
        break;
    default:
        puts("Unknown");
        break;
    }

EXIT_FUNC:

    LocalFree(pSD);

    return 0;
}
