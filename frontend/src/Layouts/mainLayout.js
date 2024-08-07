import React, { useState } from 'react';
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    DownloadOutlined,
  } from '@ant-design/icons';
import { Button, Layout, Menu } from 'antd';
import { useNavigate } from 'react-router-dom';
const { Header, Sider, Content } = Layout;


const MainLayout = (props) => {
    
    const bgColor = '#f0f0f0';
    const [collapsed, setCollapsed] = useState(false);
    const navigate = useNavigate();

    return (
        <Layout>
            <Sider trigger={null} collapsible collapsed={collapsed}>
                <Menu
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={['group1']}
                    onClick={(info) => {
                        navigate(info.key)          
                    }}
                    items={[
                        {
                            key: 'group1',
                            icon: <DownloadOutlined />,
                            label: 'Загрузка',
                            children: [
                                {
                                    key: 'personal_accounts',
                                    label: "Лицевые счета"
                                }
                            ]
                        }
                    ]}
                />
            </Sider>
            <Layout>
                <Header style={{ padding: 0, background: bgColor }}>
                    <Button
                        type="text"
                        icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                        onClick={() => setCollapsed(!collapsed)}
                        style={{
                            fontSize: '16px',
                            width: 64,
                            height: 64,
                        }}
                    />
                </Header>
                <Content
                    style={{
                        margin: '24px 16px',
                        padding: 24,
                        minHeight: '100vh',
                        background: bgColor,
                        borderRadius: 12,
                    }}
                >
                    {props.children}
                </Content>
            </Layout>
        </Layout>
    )
}

export default MainLayout;