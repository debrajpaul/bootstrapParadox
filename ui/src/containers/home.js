import React from 'react';

import { Layout, Menu, Icon,Upload, message, Button } from 'antd';

const { Header, Content, Footer, Sider } = Layout;

function Home (){
    const props = {
        name: 'file',
        action: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
        headers: {
          authorization: 'authorization-text',
        },
        onChange(info) {
          if (info.file.status !== 'uploading') {
            console.log(info.file, info.fileList);
          }
          if (info.file.status === 'done') {
            message.success(`${info.file.name} file uploaded successfully`);
          } else if (info.file.status === 'error') {
            message.error(`${info.file.name} file upload failed.`);
          }
        },
      };
    return(
        <Layout style={{width:'100%'}}>
            <Sider
            breakpoint="lg"
            collapsedWidth="0"
            onBreakpoint={broken => {
                console.log(broken);
            }}
            onCollapse={(collapsed, type) => {
                console.log(collapsed, type);
            }}
            >
            <div className="logo" style={{marginTop:60}} />
            <Menu theme="dark" mode="inline" defaultSelectedKeys={['4']}>
                <Menu.Item key="1">
                    <Icon type="user" />
                    <span className="nav-text">nav 1</span>
                </Menu.Item>
                <Menu.Item key="2">
                    <Icon type="video-camera" />
                    <span className="nav-text">nav 2</span>
                </Menu.Item>
            </Menu>
            </Sider>
            <Layout>
            <Header style={{ background: '#fff', padding: 0 }} />
            <Content style={{ margin: '24px 16px 0',width:'100%',backgroundColor:'#fff' }}>
                <div style={{position:'relative'}}>

                h3
                <div style={{position:'absolute',top:-30,right:20}}>
                    <Upload {...props}>
                            <Button type="primary">
                            <Icon type="upload" /> Click to Upload
                            </Button>
                    </Upload>
                </div>
                   

                    <div style={{display:'flex',justifyContent:'center',alignItems:'center'}}>
                        
                     
                    </div>
               </div>
            </Content>
            {/* <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer>*/}
            </Layout> 
        </Layout>

);
}


export default Home;