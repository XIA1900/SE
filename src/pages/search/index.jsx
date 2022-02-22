import { LikeOutlined, LoadingOutlined, MessageOutlined, StarOutlined } from '@ant-design/icons';
import { Button, Card, Col, Form, List, Row, Select, Tag } from 'antd';
import React from 'react';
import { useRequest, history } from 'umi';
import ArticleListContent from '@/pages/homepage/components/ArticleListContent';
import StandardFormRow from '@/pages/homepage/components/StandardFormRow';
import TagSelect from '@/pages/homepage/components/TagSelect';
import { search } from '@/services/search';
import styles from './style.less';
import { wrapConstructor } from 'lodash-decorators/utils';


const query = history.location.search;
const values = query.substr(1);
console.log(values);

const searchResults = () => {
    const [form] = Form.useForm();
    const { data, reload, loading, loadMore, loadingMore } = useRequest(
      () => {
        return search({
          values: values,
        });
      },
      {
        loadMore: true,
      },
    );
    const list = data?.list || [];
  
    const IconText = ({ type, text }) => {
      switch (type) {
        case 'star-o':
          return (
            <span>
              <StarOutlined
                style={{
                  marginRight: 8,
                }}
              />
              {text}
            </span>
          );
  
        case 'like-o':
          return (
            <span>
              <LikeOutlined
                style={{
                  marginRight: 8,
                }}
              />
              {text}
            </span>
          );
  
        case 'message':
          return (
            <span>
              <MessageOutlined
                style={{
                  marginRight: 8,
                }}
              />
              {text}
            </span>
          );
  
        default:
          return null;
      }
    };
  
    const formItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
        },
        sm: {
          span: 24,
        },
        md: {
          span: 12,
        },
      },
    };
  
    const loadMoreDom = list.length > 0 && (
      <div
        style={{
          textAlign: 'center',
          marginTop: 16,
        }}
      >
        <Button
          onClick={loadMore}
          style={{
            paddingLeft: 48,
            paddingRight: 48,
          }}
        >
          {loadingMore ? (
            <span>
              <LoadingOutlined /> Loading...
            </span>
          ) : (
            'Load More'
          )}
        </Button>
      </div>
    );
  
    return (
      <>
        <Card
          style={{
            marginTop: 24,
          }}
          bordered={false}
          bodyStyle={{
            padding: '8px 32px 32px 32px',
          }}
        >
          <List
            size="large"
            loading={loading}
            rowKey="id"
            itemLayout="vertical"
            loadMore={loadMoreDom}
            dataSource={list}
            renderItem={(item) => (
              <List.Item
                key={item.id}
                actions={[
                  <IconText key="collection" type="star-o" text={item.star} />,
                  <IconText key="like" type="like-o" text={item.like} />,
                  <IconText key="reply" type="message" text={item.message} />,
                ]}
                extra={<div className={styles.listItemExtra} />}
              >
                <List.Item.Meta
                  title={
                    <a className={styles.listItemMetaTitle} href={item.href}>
                      {item.title}
                    </a>
                  }
                />
                <ArticleListContent data={item} />
              </List.Item>
            )}
          />
        </Card>
      </>
    );
  };
  
  export default searchResults;
