module.exports = {
  apps : [
    {
      name: 'api-gateway',
      script: 'air',
      cwd: 'go-grpc-api-gateway',    
    },
    {
      name: 'order-svc',
      script: 'air',
      cwd: 'go-grpc-order-svc',    
    },
    {
      name: 'auth-svc',
      script: 'air',
      cwd: 'go-grpc-auth-svc',    
    },
    {
      name: 'product-svc',
      script: 'air',
      cwd: 'go-grpc-product-svc',    
    },
  ],
};
