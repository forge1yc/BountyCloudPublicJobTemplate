# 设置基本的镜像，后续命令都以这个镜像为基础
FROM centos:centos7
# 作者信息
MAINTAINER forge1yc, houyichaochao@gmail.com
# RUN命令会在上面指定的镜像里执行任何命令
RUN yum install passwd openssl openssh-server -y
RUN echo '123456' | passwd --stdin root
RUN ssh-keygen -q -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key -N ''
RUN ssh-keygen -q -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N ''
RUN sed -i '/^session\s\+required\s\+pam_loginuid.so/s/^/#/' /etc/pam.d/sshd
RUN mkdir -p /root/.ssh && chown root.root /root && chmod 700 /root/.ssh
# 暴露ssh端口22
EXPOSE 22
# 设定运行镜像时的默认命令：输出ip，并以daemon方式启动sshd
CMD ip addr ls eth0 | awk '{print $2}' | egrep -o '([0-9]+\.){3}[0-9]+';/usr/sbin/sshd -D