import React,{useState} from 'react';

import { Layout, Menu, Icon,Upload, message, Button,Input } from 'antd';
 const {Search}=Input;
const { Header, Content, Footer, Sider } = Layout;

function Home (){
  const [state, setstate] = useState([]);
 const search=async value=>{
     if(value.length>3){
      let res=  await fetch('url',{method:'post',headers:{'Content-Type':'application/json'},body:JSON.stringify({})})
        let reques=await res.json();
         
     }
  }
   


    const props = {
        name: 'file',
        action: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
        headers: {
          'Content-Type': 'multipart/form-data'
        },
        onChange(info) {
          if (info.file.status !== 'uploading') {
           // console.log(info.file, info.fileList);
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
            }}
            onCollapse={(collapsed, type) => {
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
            <div style={{width:'60%',justifyContent:'center',alignItems:'center',alignSelf:'center',marginTop:16}}>
              <Search placeholder="input search text" onSearch={search}  size="large" enterButton />
            </div>
           
            <Content style={{ margin: '24px 16px 0',width:'100%',backgroundColor:'#fff' }}>
           
                <div style={{position:'relative'}}>

                
                <div style={{position:'absolute',top:-60,right:30}}>
                    <Upload {...props}>
                            <Button type="primary">
                            <Icon type="upload" /> Upload Here!
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