import axios from 'axios';

const AbstractRepository = axios.create({
  baseURL: '/api',
  timeout: 10000,
  params: {},
  headers: {},
  withCredentials: true,
});

export default AbstractRepository;
