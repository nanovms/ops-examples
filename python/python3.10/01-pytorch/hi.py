import torch
import torch.nn as nn
import torch.optim as optim

class Net(nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.hidden_layer = nn.Linear(1, 1)
        self.hidden_layer.weight =
torch.nn.Parameter(torch.tensor([[1.58]]))
        self.hidden_layer.bias =
torch.nn.Parameter(torch.tensor([-0.14]))

        self.output_layer = nn.Linear(1, 1)
        self.output_layer.weight =
torch.nn.Parameter(torch.tensor([[2.45]]))
        self.output_layer.bias =
torch.nn.Parameter(torch.tensor([-0.11]))

    def forward(self, x):
        x = torch.sigmoid(self.hidden_layer(x))
        x = torch.sigmoid(self.output_layer(x))
        return x


net = Net()
print(f"network topology: {net}")

print(f"w_l1 = {round(net.hidden_layer.weight.item(), 4)}")
print(f"b_l1 = {round(net.hidden_layer.bias.item(), 4)}")
print(f"w_l2 = {round(net.output_layer.weight.item(), 4)}")
print(f"b_l2 = {round(net.output_layer.bias.item(), 4)}")

# run input data forward through network
input_data = torch.tensor([0.8])
output = net(input_data)
print(f"a_l2 = {round(output.item(), 4)}")

# backpropagate gradient
target = torch.tensor([1.])
criterion = nn.MSELoss()
loss = criterion(output, target)
net.zero_grad()
loss.backward()

# update weights and biases
optimizer = optim.SGD(net.parameters(), lr=0.1)
optimizer.step()

print(f"updated_w_l1 = {round(net.hidden_layer.weight.item(), 4)}")
print(f"updated_b_l1 = {round(net.hidden_layer.bias.item(), 4)}")
print(f"updated_w_l2 = {round(net.output_layer.weight.item(), 4)}")
print(f"updated_b_l2 = {round(net.output_layer.bias.item(), 4)}")

output = net(input_data)
print(f"updated_a_l2 = {round(output.item(), 4)}")
